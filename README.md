# Glide Node.js Plugin

External Node.js plugin for Glide - provides Node.js and package manager integration.

## Overview

This plugin provides Node.js functionality for Glide, including:

- Automatic package manager detection (npm, yarn, pnpm, bun)
- Framework detection (React, Vue, Angular, Next.js, etc.)
- Package.json integration and script running
- Workspace/monorepo support
- Project metadata extraction

## Installation

### Method 1: Build from Source

```bash
# Clone the repository
git clone https://github.com/ivannovak/glide-plugin-node
cd glide-plugin-node

# Build the plugin
make build

# Install to PATH
sudo cp glide-plugin-node /usr/local/bin/
```

### Method 2: Go Install (when published)

```bash
go install github.com/ivannovak/glide-plugin-node/cmd/glide-plugin-node@latest
```

## Usage

Once installed, the plugin provides Node.js commands to Glide:

```bash
# Install dependencies (auto-detects package manager)
glide install

# Run package.json scripts
glide run test
glide run build
glide run dev
glide run lint -- --fix
```

## Commands

### `install` (alias: `i`)

Install Node.js dependencies using the detected package manager.

```bash
glide install
```

### `run <script> [args...]`

Run any script defined in your `package.json`:

```bash
glide run test
glide run build
glide run dev
glide run custom-script -- --with-args
```

## Detection

The plugin automatically detects:

- **Package Manager**: Based on lock files (yarn.lock, pnpm-lock.yaml, bun.lockb, package-lock.json)
- **Frameworks**: React, Vue, Angular, Next.js, Nuxt, Express, NestJS, TypeScript, Svelte, Remix, Gatsby, Vite
- **Module Type**: ESM vs CommonJS
- **Workspaces**: Monorepo detection
- **Node Version**: From `engines` field in package.json

## Development

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all

# Run tests
make test

# Tidy dependencies
make tidy
```

### Current Status

**Phase 2: Node Plugin Creation** (In Progress)

The plugin structure is complete, but currently relies on Glide's internal packages via a local replace directive. To make this fully standalone:

1. Glide core needs to implement `sdk.RunPlugin()` function
2. SDK needs to expose extension data access in commands
3. Remove the local replace directive

This will be completed in a future phase when the public SDK API is finalized.

### Project Structure

```
glide-plugin-node/
├── cmd/
│   └── glide-plugin-node/
│       └── main.go              # Plugin entry point
├── internal/
│   ├── commands/
│   │   └── node.go              # Node.js commands
│   └── plugin/
│       ├── detector.go          # Node.js project detection
│       └── plugin.go            # Plugin implementation
├── Makefile                     # Build automation
└── README.md                    # This file
```

## License

MIT
