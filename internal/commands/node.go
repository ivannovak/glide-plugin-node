package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ivannovak/glide/pkg/plugin/sdk"
	sdkv1 "github.com/ivannovak/glide/pkg/plugin/sdk/v1"
	"github.com/spf13/cobra"
)

// NewNodeCommands returns all Node.js commands
func NewNodeCommands() []*sdk.PluginCommandDefinition {
	return []*sdk.PluginCommandDefinition{
		NewInstallCommand(),
		NewRunCommand(),
	}
}

// NewInstallCommand creates the 'install' command
func NewInstallCommand() *sdk.PluginCommandDefinition {
	return &sdk.PluginCommandDefinition{
		Name:  "install",
		Use:   "install",
		Short: "Install Node.js dependencies",
		Long: `Install Node.js dependencies using the detected package manager.

This command automatically detects whether you're using npm, yarn, pnpm, or bun
and runs the appropriate install command.

Examples:
  glide install              # Install all dependencies
`,
		Aliases: []string{"i"},
		RunE:    executeInstall,
	}
}

// NewRunCommand creates the 'run' command
func NewRunCommand() *sdk.PluginCommandDefinition {
	return &sdk.PluginCommandDefinition{
		Name:  "run",
		Use:   "run <script> [args...]",
		Short: "Run a package.json script",
		Long: `Run any script defined in your package.json file.

This command uses the detected package manager (npm, yarn, pnpm, or bun)
to execute the specified script.

Examples:
  glide run test             # Run the test script
  glide run build            # Run the build script
  glide run dev              # Run the dev script
  glide run lint -- --fix    # Run lint script with additional args
`,
		Args: cobra.MinimumNArgs(1),
		RunE: executeRun,
	}
}

// executeInstall runs the install command
func executeInstall(cmd *cobra.Command, args []string) error {
	// Get project context
	ctx := getProjectContext(cmd)
	if ctx == nil {
		return fmt.Errorf("project context not available")
	}

	// Get package manager from Node extension
	packageManager := getPackageManager(ctx)

	// Build install command
	var installCmd *exec.Cmd
	switch packageManager {
	case "yarn":
		installCmd = exec.Command("yarn", "install")
	case "pnpm":
		installCmd = exec.Command("pnpm", "install")
	case "bun":
		installCmd = exec.Command("bun", "install")
	default: // npm
		installCmd = exec.Command("npm", "install")
	}

	// Set working directory
	installCmd.Dir = ctx.Root
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	installCmd.Stdin = os.Stdin

	fmt.Printf("Installing dependencies with %s...\n", packageManager)
	return installCmd.Run()
}

// executeRun runs a package.json script
func executeRun(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("script name required")
	}

	scriptName := args[0]
	scriptArgs := args[1:]

	// Get project context
	ctx := getProjectContext(cmd)
	if ctx == nil {
		return fmt.Errorf("project context not available")
	}

	// Get package manager from Node extension
	packageManager := getPackageManager(ctx)

	// Build run command
	var runCmd *exec.Cmd
	switch packageManager {
	case "yarn":
		runCmd = exec.Command("yarn", append([]string{scriptName}, scriptArgs...)...)
	case "pnpm":
		runCmd = exec.Command("pnpm", append([]string{scriptName}, scriptArgs...)...)
	case "bun":
		runCmd = exec.Command("bun", append([]string{"run", scriptName}, scriptArgs...)...)
	default: // npm
		runCmd = exec.Command("npm", append([]string{"run", scriptName}, scriptArgs...)...)
	}

	// Set working directory
	runCmd.Dir = ctx.Root
	runCmd.Stdout = os.Stdout
	runCmd.Stderr = os.Stderr
	runCmd.Stdin = os.Stdin

	fmt.Printf("Running script '%s' with %s...\n", scriptName, packageManager)
	return runCmd.Run()
}

// getProjectContext extracts the project context from the command
func getProjectContext(cmd *cobra.Command) *sdkv1.ProjectContext {
	ctxValue := cmd.Context().Value("project_context")
	if ctxValue == nil {
		return nil
	}

	ctx, ok := ctxValue.(*sdkv1.ProjectContext)
	if !ok {
		return nil
	}

	return ctx
}

// getPackageManager gets the package manager from context
// For now, it tries to detect from the project root
func getPackageManager(ctx *sdkv1.ProjectContext) string {
	// TODO: Get from extensions when SDK provides access
	// For now, default to npm
	return "npm"
}
