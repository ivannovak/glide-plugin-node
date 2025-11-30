package plugin

import (
	"context"

	"github.com/ivannovak/glide-plugin-node/pkg/version"
	"github.com/ivannovak/glide/v3/pkg/plugin/sdk/v2"
)

// Config defines the plugin's type-safe configuration.
// Users configure this in .glide.yml under plugins.node
type Config struct {
	// PreferYarn uses yarn instead of npm when available
	PreferYarn bool `json:"preferYarn" yaml:"preferYarn"`

	// PreferPnpm uses pnpm instead of npm when available
	PreferPnpm bool `json:"preferPnpm" yaml:"preferPnpm"`

	// EnableTypeScript enables TypeScript-specific detection
	EnableTypeScript bool `json:"enableTypeScript" yaml:"enableTypeScript"`
}

// DefaultConfig returns sensible defaults
func DefaultConfig() Config {
	return Config{
		PreferYarn:       false,
		PreferPnpm:       false,
		EnableTypeScript: true,
	}
}

// NodePlugin implements the SDK v2 Plugin interface for Node.js detection
type NodePlugin struct {
	v2.BasePlugin[Config]
}

// New creates a new Node.js plugin instance
func New() *NodePlugin {
	return &NodePlugin{}
}

// Metadata returns plugin information
func (p *NodePlugin) Metadata() v2.Metadata {
	return v2.Metadata{
		Name:        "node",
		Version:     version.Version,
		Author:      "Glide Team",
		Description: "Node.js framework detector for Glide",
		License:     "MIT",
		Homepage:    "https://github.com/ivannovak/glide-plugin-node",
		Tags:        []string{"language", "node", "nodejs", "javascript", "typescript", "detector"},
	}
}

// Configure is called with the type-safe configuration
func (p *NodePlugin) Configure(ctx context.Context, config Config) error {
	return p.BasePlugin.Configure(ctx, config)
}

// Commands returns the list of commands this plugin provides.
// Note: This is a framework detector plugin, so it doesn't provide CLI commands.
func (p *NodePlugin) Commands() []v2.Command {
	return []v2.Command{}
}

// Init is called once after plugin load
func (p *NodePlugin) Init(ctx context.Context) error {
	return nil
}

// HealthCheck returns nil if the plugin is healthy
func (p *NodePlugin) HealthCheck(ctx context.Context) error {
	return nil
}
