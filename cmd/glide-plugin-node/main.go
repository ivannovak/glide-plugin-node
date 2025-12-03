package main

import (
	"fmt"
	"os"

	"github.com/glide-cli/glide-plugin-node/internal/plugin"
	"github.com/glide-cli/glide/v3/pkg/plugin/sdk/v2"
)

func main() {
	// Initialize the Node.js plugin
	nodePlugin := plugin.New()

	// Run the plugin using SDK v2
	if err := v2.Serve(nodePlugin); err != nil {
		fmt.Fprintf(os.Stderr, "Plugin error: %v\n", err)
		os.Exit(1)
	}
}
