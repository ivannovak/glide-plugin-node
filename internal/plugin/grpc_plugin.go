package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	v1 "github.com/ivannovak/glide/pkg/plugin/sdk/v1"
)

// GRPCPlugin implements the gRPC GlidePluginServer interface
type GRPCPlugin struct {
	*v1.BasePlugin
	detector *NodeDetector
}

// NewGRPCPlugin creates a new gRPC-based Node.js plugin
func NewGRPCPlugin() *GRPCPlugin {
	metadata := &v1.PluginMetadata{
		Name:        "node",
		Version:     "1.0.0",
		Author:      "Glide Team",
		Description: "Node.js and package manager integration for Glide",
		Homepage:    "https://github.com/ivannovak/glide-plugin-node",
		License:     "MIT",
		Tags:        []string{"language", "node", "nodejs", "javascript", "typescript"},
		Aliases:     []string{"nodejs", "js", "ts"},
		Namespaced:  false,
	}

	p := &GRPCPlugin{
		BasePlugin: v1.NewBasePlugin(metadata),
		detector:   NewNodeDetector(),
	}

	// Register all Node commands
	p.registerCommands()

	return p
}

// registerCommands registers all Node.js-related commands
func (p *GRPCPlugin) registerCommands() {
	// Install command
	installHandler := v1.NewSimpleCommand(
		&v1.CommandInfo{
			Name:        "install",
			Description: "Install Node.js dependencies",
			Category:    "dependencies",
			Aliases:     []string{"i"},
			Visibility:  "project-only",
		},
		func(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
			return p.executeInstall(ctx, req)
		},
	)
	p.RegisterCommand("install", installHandler)

	// Run command
	runHandler := v1.NewSimpleCommand(
		&v1.CommandInfo{
			Name:        "run",
			Description: "Run a package.json script",
			Category:    "run",
			Visibility:  "project-only",
		},
		func(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
			return p.executeRun(ctx, req)
		},
	)
	p.RegisterCommand("run", runHandler)
}

// executeInstall runs the install command
func (p *GRPCPlugin) executeInstall(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
	workDir := req.WorkDir
	if workDir == "" {
		workDir = "."
	}

	// Detect package manager
	packageManager := p.detector.detectPackageManager(workDir)

	// Build install command
	var cmdParts []string
	switch packageManager {
	case "yarn":
		cmdParts = []string{"yarn", "install"}
	case "pnpm":
		cmdParts = []string{"pnpm", "install"}
	case "bun":
		cmdParts = []string{"bun", "install"}
	default: // npm
		cmdParts = []string{"npm", "install"}
	}

	// Add any additional args
	cmdParts = append(cmdParts, req.Args...)

	// Execute
	cmd := exec.CommandContext(ctx, cmdParts[0], cmdParts[1:]...)
	cmd.Dir = workDir

	// Set environment - start with parent environment
	cmd.Env = os.Environ()
	// Override/add custom environment variables
	for k, v := range req.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	output, err := cmd.CombinedOutput()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return &v1.ExecuteResponse{
				Success:  false,
				ExitCode: 1,
				Error:    err.Error(),
			}, nil
		}
	}

	return &v1.ExecuteResponse{
		Success:  exitCode == 0,
		ExitCode: int32(exitCode),
		Stdout:   output,
	}, nil
}

// executeRun runs a package.json script
func (p *GRPCPlugin) executeRun(ctx context.Context, req *v1.ExecuteRequest) (*v1.ExecuteResponse, error) {
	if len(req.Args) == 0 {
		return &v1.ExecuteResponse{
			Success:  false,
			ExitCode: 1,
			Error:    "script name required",
		}, nil
	}

	scriptName := req.Args[0]
	scriptArgs := req.Args[1:]

	workDir := req.WorkDir
	if workDir == "" {
		workDir = "."
	}

	// Detect package manager
	packageManager := p.detector.detectPackageManager(workDir)

	// Build run command
	var cmdParts []string
	switch packageManager {
	case "yarn":
		cmdParts = append([]string{"yarn", scriptName}, scriptArgs...)
	case "pnpm":
		cmdParts = append([]string{"pnpm", scriptName}, scriptArgs...)
	case "bun":
		cmdParts = append([]string{"bun", "run", scriptName}, scriptArgs...)
	default: // npm
		cmdParts = append([]string{"npm", "run", scriptName}, scriptArgs...)
	}

	// Execute
	cmd := exec.CommandContext(ctx, cmdParts[0], cmdParts[1:]...)
	cmd.Dir = workDir

	// Set environment - start with parent environment
	cmd.Env = os.Environ()
	// Override/add custom environment variables
	for k, v := range req.Env {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", k, v))
	}

	output, err := cmd.CombinedOutput()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			return &v1.ExecuteResponse{
				Success:  false,
				ExitCode: 1,
				Error:    err.Error(),
			}, nil
		}
	}

	return &v1.ExecuteResponse{
		Success:  exitCode == 0,
		ExitCode: int32(exitCode),
		Stdout:   output,
	}, nil
}

// DetectContext implements context detection for Node.js projects
func (p *GRPCPlugin) DetectContext(ctx context.Context, req *v1.ContextRequest) (*v1.ContextResponse, error) {
	// Use the detector to check if this is a Node.js project
	projectRoot := req.ProjectRoot
	if projectRoot == "" {
		projectRoot = req.WorkingDir
	}

	// Check if package.json exists
	packageJSONPath := filepath.Join(projectRoot, "package.json")
	if _, err := os.Stat(packageJSONPath); os.IsNotExist(err) {
		// Not a Node.js project
		return &v1.ContextResponse{
			ExtensionName: "node",
			Detected:      false,
		}, nil
	}

	// Run detection
	data, err := p.detector.Detect(ctx, projectRoot)
	if err != nil || data == nil {
		return &v1.ContextResponse{
			ExtensionName: "node",
			Detected:      false,
		}, nil
	}

	// Convert data map to response
	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return &v1.ContextResponse{
			ExtensionName: "node",
			Detected:      false,
		}, nil
	}

	detected, _ := dataMap["node_detected"].(bool)
	if !detected {
		return &v1.ContextResponse{
			ExtensionName: "node",
			Detected:      false,
		}, nil
	}

	// Build response
	resp := &v1.ContextResponse{
		ExtensionName: "node",
		Detected:      true,
		Metadata:      make(map[string]string),
		Frameworks:    []string{},
		Tools:         []string{},
	}

	// Convert metadata
	for k, v := range dataMap {
		switch k {
		case "node_detected", "scripts", "frameworks":
			continue
		case "package_manager":
			if str, ok := v.(string); ok {
				resp.Metadata[k] = str
				resp.Tools = append(resp.Tools, str)
			}
		default:
			if str, ok := v.(string); ok {
				resp.Metadata[k] = str
			} else if b, ok := v.(bool); ok {
				resp.Metadata[k] = fmt.Sprintf("%v", b)
			}
		}
	}

	// Extract version
	if nodeVersion, ok := dataMap["node_version"].(string); ok {
		resp.Version = nodeVersion
	}

	// Extract frameworks
	if frameworks, ok := dataMap["frameworks"].([]string); ok {
		resp.Frameworks = frameworks
	}

	// Convert module type to framework indicator
	if moduleType, ok := dataMap["module_type"].(string); ok {
		if moduleType == "module" {
			resp.Frameworks = append(resp.Frameworks, "ES Modules")
		}
	}

	return resp, nil
}
