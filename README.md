# glide-plugin-node

[![CI](https://github.com/ivannovak/glide-plugin-node/actions/workflows/ci.yml/badge.svg)](https://github.com/ivannovak/glide-plugin-node/actions/workflows/ci.yml)
[![Semantic Release](https://github.com/ivannovak/glide-plugin-node/actions/workflows/semantic-release.yml/badge.svg)](https://github.com/ivannovak/glide-plugin-node/actions/workflows/semantic-release.yml)

Node.js and package manager integration plugin for [Glide CLI](https://github.com/ivannovak/glide).

## Overview

This plugin provides Node.js project detection and intelligent package manager support for Glide. When installed, Glide will automatically detect Node.js projects and provide smart command completions based on your package manager (npm, yarn, pnpm, or bun).

## Installation

### From GitHub Releases (Recommended)

```bash
glide plugins install github.com/ivannovak/glide-plugin-node
```

### From Source

```bash
# Clone the repository
git clone https://github.com/ivannovak/glide-plugin-node.git
cd glide-plugin-node

# Build and install (requires Go 1.24+)
make install
```

## What It Detects

The plugin automatically detects Node.js projects by looking for:

- **Required files**: `package.json`
- **Lock files**: `package-lock.json`, `yarn.lock`, `pnpm-lock.yaml`, `bun.lockb`
- **Config files**: `.npmrc`, `.yarnrc`, `pnpm-workspace.yaml`
- **Directories**: `node_modules/`

### Package Manager Auto-Detection

The plugin intelligently detects your preferred package manager:

1. **bun** - If `bun.lockb` exists
2. **pnpm** - If `pnpm-lock.yaml` exists
3. **yarn** - If `yarn.lock` exists
4. **npm** - Default fallback

### Enhanced Detection

The plugin extracts additional metadata from your Node.js project:

- **Node version**: Detected from `package.json` engines field
- **Package manager**: Auto-detected from lock files
- **Module type**: Detects ES Modules vs CommonJS
- **Frameworks**: Recognizes React, Vue, Angular, Next.js, Nuxt, Express, NestJS, Svelte, Remix, Gatsby, Vite
- **Scripts**: Lists available npm scripts
- **Workspaces**: Detects monorepo configurations

## Available Commands

Once a Node.js project is detected, the following commands become available:

### Package Management
- `install` (alias: `i`) - Install dependencies with detected package manager
  - With args: Installs specified packages
  - Without args: Installs all dependencies

### Script Execution
- `run <script>` - Run a package.json script with detected package manager
  - Automatically uses the correct package manager
  - Forwards all arguments to the script

## Configuration

The plugin works out-of-the-box without configuration. However, you can customize behavior in your `.glide.yml`:

```yaml
plugins:
  node:
    enabled: true
    # Additional configuration options can be added here in the future
```

## Examples

### Basic Node.js Project

```bash
# Navigate to your Node.js project
cd my-app

# Glide automatically detects Node and package manager
glide help

# Install dependencies (uses detected package manager)
glide install

# Install a specific package
glide install lodash

# Run package.json scripts
glide run dev
glide run build
glide run test
```

### Package Manager Detection

```bash
# With yarn.lock - uses yarn
cd my-yarn-project
glide install  # Executes: yarn install
glide run dev  # Executes: yarn dev

# With pnpm-lock.yaml - uses pnpm
cd my-pnpm-project
glide install  # Executes: pnpm install
glide run dev  # Executes: pnpm dev

# With bun.lockb - uses bun
cd my-bun-project
glide install  # Executes: bun install
glide run dev  # Executes: bun run dev
```

### Common Workflows

```bash
# Development workflow
glide install     # Install dependencies
glide run dev     # Start development server
glide run build   # Build for production
glide run test    # Run tests

# Managing dependencies
glide install react  # Add dependency
glide run lint       # Run linter
glide run format     # Format code
```

## Development

### Prerequisites

- Go 1.24 or higher
- Make (optional, for convenience targets)

### Building

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage

# Run linters
make lint

# Format code
make fmt
```

### Testing

The plugin includes comprehensive tests for:

- Node version detection
- Package manager detection
- Module type detection
- Framework detection
- Command execution

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

## Contributing

Contributions are welcome! Please:

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass (`make test`)
6. Submit a pull request

## License

MIT License - see [LICENSE](LICENSE) for details.

## Related Projects

- [Glide](https://github.com/ivannovak/glide) - The main Glide CLI
- [glide-plugin-go](https://github.com/ivannovak/glide-plugin-go) - Go plugin for Glide
- [glide-plugin-php](https://github.com/ivannovak/glide-plugin-php) - PHP plugin for Glide
- [glide-plugin-docker](https://github.com/ivannovak/glide-plugin-docker) - Docker plugin for Glide

## Support

- [GitHub Issues](https://github.com/ivannovak/glide-plugin-node/issues)
- [Glide Documentation](https://github.com/ivannovak/glide#readme)
- [Plugin Development Guide](https://github.com/ivannovak/glide/blob/main/docs/PLUGIN_DEVELOPMENT.md)
