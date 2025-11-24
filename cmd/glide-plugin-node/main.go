package main

import (
	"os"

	"github.com/ivannovak/glide-plugin-node/internal/plugin"
	sdk "github.com/ivannovak/glide/pkg/plugin/sdk/v1"
)

func main() {
	// Initialize the Node.js gRPC plugin
	nodePlugin := plugin.NewGRPCPlugin()

	// Run the plugin using the SDK
	if err := sdk.RunPlugin(nodePlugin); err != nil {
		os.Exit(1)
	}
}
