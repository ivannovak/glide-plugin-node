package plugin

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
)

// NodeDetector implements the SDK ContextExtension interface for Node.js detection
type NodeDetector struct{}

// PackageJSON represents the structure of package.json
type PackageJSON struct {
	Name            string            `json:"name"`
	Version         string            `json:"version"`
	Scripts         map[string]string `json:"scripts"`
	Dependencies    map[string]string `json:"dependencies"`
	DevDependencies map[string]string `json:"devDependencies"`
	Engines         map[string]string `json:"engines"`
	Type            string            `json:"type"`
	Main            string            `json:"main"`
	Private         bool              `json:"private"`
	Workspaces      interface{}       `json:"workspaces"`
}

// NewNodeDetector creates a new Node.js detector
func NewNodeDetector() *NodeDetector {
	return &NodeDetector{}
}

// Name returns the unique identifier for this extension
func (d *NodeDetector) Name() string {
	return "node"
}

// Detect analyzes the project environment and returns Node.js-specific context data
func (d *NodeDetector) Detect(ctx context.Context, projectRoot string) (interface{}, error) {
	// Check if package.json exists
	packageJSONPath := filepath.Join(projectRoot, "package.json")
	if _, err := os.Stat(packageJSONPath); os.IsNotExist(err) {
		// No package.json, not a Node project
		return nil, nil
	}

	// Read and parse package.json
	pkg, err := d.readPackageJSON(packageJSONPath)
	if err != nil {
		// package.json exists but can't be read/parsed
		return map[string]interface{}{
			"node_detected": true,
			"error":         "failed to parse package.json",
		}, nil
	}

	// Detect package manager
	packageManager := d.detectPackageManager(projectRoot)

	// Detect frameworks
	frameworks := d.detectFrameworks(pkg)

	// Build the extension data structure
	result := map[string]interface{}{
		"node_detected":   true,
		"package_manager": packageManager,
		"project_name":    pkg.Name,
		"project_version": pkg.Version,
		"scripts":         pkg.Scripts,
		"frameworks":      frameworks,
	}

	// Add optional fields
	if pkg.Type != "" {
		result["module_type"] = pkg.Type
	}
	if engines, ok := pkg.Engines["node"]; ok {
		result["node_version"] = engines
	}
	if pkg.Private {
		result["private"] = true
	}
	if pkg.Workspaces != nil {
		result["workspaces"] = true
		result["monorepo"] = true
	}

	return result, nil
}

// Merge merges two context data structures
func (d *NodeDetector) Merge(existing, new interface{}) (interface{}, error) {
	// If either is nil, return the non-nil one
	if existing == nil {
		return new, nil
	}
	if new == nil {
		return existing, nil
	}

	// Type assert both to maps
	existingMap, ok1 := existing.(map[string]interface{})
	newMap, ok2 := new.(map[string]interface{})

	if !ok1 || !ok2 {
		// If either is not a map, prefer new
		return new, nil
	}

	// Merge maps (new values override existing)
	result := make(map[string]interface{})
	for k, v := range existingMap {
		result[k] = v
	}
	for k, v := range newMap {
		result[k] = v
	}

	return result, nil
}

// readPackageJSON reads and parses package.json
func (d *NodeDetector) readPackageJSON(path string) (*PackageJSON, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var pkg PackageJSON
	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, err
	}

	return &pkg, nil
}

// detectPackageManager detects which package manager is being used
func (d *NodeDetector) detectPackageManager(projectRoot string) string {
	// Check for lock files to determine package manager
	lockFiles := map[string]string{
		"yarn.lock":        "yarn",
		"pnpm-lock.yaml":   "pnpm",
		"bun.lockb":        "bun",
		"package-lock.json": "npm",
	}

	for lockFile, pm := range lockFiles {
		if _, err := os.Stat(filepath.Join(projectRoot, lockFile)); err == nil {
			return pm
		}
	}

	// Default to npm
	return "npm"
}

// detectFrameworks detects specific Node.js frameworks from dependencies
func (d *NodeDetector) detectFrameworks(pkg *PackageJSON) []string {
	var frameworks []string

	// Combine all dependencies
	allDeps := make(map[string]bool)
	for dep := range pkg.Dependencies {
		allDeps[dep] = true
	}
	for dep := range pkg.DevDependencies {
		allDeps[dep] = true
	}

	// Check for popular frameworks
	frameworkChecks := map[string]string{
		"react":         "React",
		"next":          "Next.js",
		"vue":           "Vue",
		"nuxt":          "Nuxt",
		"@angular/core": "Angular",
		"express":       "Express",
		"@nestjs/core":  "NestJS",
		"typescript":    "TypeScript",
		"svelte":        "Svelte",
		"@remix-run/react": "Remix",
		"gatsby":        "Gatsby",
		"vite":          "Vite",
	}

	for dep, name := range frameworkChecks {
		if allDeps[dep] {
			frameworks = append(frameworks, name)
		}
	}

	return frameworks
}
