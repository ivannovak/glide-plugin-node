package plugin

import (
	"github.com/ivannovak/glide/pkg/plugin"
	"github.com/ivannovak/glide/pkg/plugin/sdk"
	"github.com/ivannovak/glide-plugin-node/internal/commands"
	"github.com/spf13/cobra"
)

// NodePlugin implements the SDK Plugin interfaces for Node.js functionality
type NodePlugin struct {
	detector *NodeDetector
}

// New creates a new Node.js plugin instance
func New() *NodePlugin {
	return &NodePlugin{
		detector: NewNodeDetector(),
	}
}

// Name returns the plugin identifier
func (p *NodePlugin) Name() string {
	return "node"
}

// Version returns the plugin version
func (p *NodePlugin) Version() string {
	return "1.0.0"
}

// Description returns the plugin description
func (p *NodePlugin) Description() string {
	return "Node.js and package manager integration for Glide"
}

// Register adds plugin commands to the command tree
func (p *NodePlugin) Register(root *cobra.Command) error {
	// Get the command definitions from the SDK layer
	cmdDefs := p.ProvideCommands()

	// Convert and register each command with the root
	for _, cmdDef := range cmdDefs {
		if cmdDef != nil {
			cobraCmd := cmdDef.ToCobraCommand()

			// Wrap the command to inject project context
			// We need to do this because plugin commands don't have direct access to the app context
			p.wrapCommandWithContext(cobraCmd, root)

			root.AddCommand(cobraCmd)
		}
	}

	return nil
}

// wrapCommandWithContext wraps a command to inject project context from the root command
func (p *NodePlugin) wrapCommandWithContext(cmd *cobra.Command, root *cobra.Command) {
	// Store the original RunE
	originalRunE := cmd.RunE
	if originalRunE == nil {
		return
	}

	// Wrap it to inject context
	cmd.RunE = func(c *cobra.Command, args []string) error {
		// Get the root command to access its context
		// The context should be set by the main CLI before execution
		rootCtx := c.Root().Context()
		if rootCtx != nil {
			// Set the context on this command
			c.SetContext(rootCtx)
		}

		// Call the original RunE
		return originalRunE(c, args)
	}
}

// Configure allows plugin-specific configuration
func (p *NodePlugin) Configure(config map[string]interface{}) error {
	// Node plugin doesn't require specific configuration yet
	// Future: Could add default package manager, script preferences, etc.
	return nil
}

// Metadata returns plugin information
func (p *NodePlugin) Metadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name:        "node",
		Version:     "1.0.0",
		Author:      "Glide Team",
		Description: "Node.js and package manager integration for Glide",
		Aliases:     []string{},
		Commands: []plugin.CommandInfo{
			{
				Name:        "install",
				Category:    "Node.js",
				Description: "Install Node.js dependencies",
				Aliases:     []string{"i"},
			},
			{
				Name:        "run",
				Category:    "Node.js",
				Description: "Run package.json scripts",
				Aliases:     []string{},
			},
		},
		BuildTags:  []string{},
		ConfigKeys: []string{"node"},
	}
}

// ProvideContext returns the context extension for Node.js detection
func (p *NodePlugin) ProvideContext() sdk.ContextExtension {
	return p.detector
}

// ProvideCommands returns the commands provided by this plugin
func (p *NodePlugin) ProvideCommands() []*sdk.PluginCommandDefinition {
	return commands.NewNodeCommands()
}
