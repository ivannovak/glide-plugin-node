## [3.0.0](https://github.com/ivannovak/glide-plugin-node/compare/v2.3.3...v3.0.0) (2025-12-01)


### ⚠ BREAKING CHANGES

* This plugin now requires Glide v2.4.0+ with SDK v2.

Migration to SDK v2:
- Replace v1.BasePlugin with v2.BasePlugin[Config]
- Add type-safe Config struct with preferYarn, preferPnpm,
  and enableTypeScript options
- Update main.go to use v2.Serve()
- Create new plugin.go with SDK v2 patterns
- Remove legacy grpc_plugin.go (v1 implementation)

The plugin now uses the declarative SDK v2 pattern with:
- Type-safe configuration via Go generics
- Unified lifecycle management
- Declarative metadata via Metadata() method

Note: go.mod includes a replace directive pointing to local glide
repo until v2.4.0 with SDK v2 is released.

* feat!: upgrade to glide SDK v3.0.0
* Updates module dependency from glide/v2 to glide/v3 v3.0.0.
This aligns with the SDK v2 type-safe configuration system released in glide v3.0.0.

- Update go.mod to require github.com/ivannovak/glide/v3 v3.0.0
- Remove local replace directive (now using published version)
- Update all imports from /v2/ to /v3/

* ci: add CI workflow for PR validation

### Features

* upgrade to glide SDK v3.0.0 ([#1](https://github.com/ivannovak/glide-plugin-node/issues/1)) ([d10b221](https://github.com/ivannovak/glide-plugin-node/commit/d10b221702bbb83bfc827380ed8d834046bfb410))

## [2.3.3](https://github.com/ivannovak/glide-plugin-node/compare/v2.3.2...v2.3.3) (2025-11-25)


### Bug Fixes

* remove command registration from detector-only plugin ([b228cac](https://github.com/ivannovak/glide-plugin-node/commit/b228cacb5d342d6fdc5be5f4826e4929de4f3e81))

## [2.3.2](https://github.com/ivannovak/glide-plugin-node/compare/v2.3.1...v2.3.2) (2025-11-25)


### Bug Fixes

* correct build path in release workflow ([283ec0b](https://github.com/ivannovak/glide-plugin-node/commit/283ec0b372283a262e130500c7e35b01b735638d))

## [2.3.1](https://github.com/ivannovak/glide-plugin-node/compare/v2.3.0...v2.3.1) (2025-11-25)


### Bug Fixes

* remove CI dependency from release workflow ([bdd2960](https://github.com/ivannovak/glide-plugin-node/commit/bdd2960fe33459f7ad0d24c341486d27b54132ab))

## [2.3.0](https://github.com/ivannovak/glide-plugin-node/compare/v2.2.0...v2.3.0) (2025-11-25)


### Features

* use published Glide v2.2.0 ([32ece41](https://github.com/ivannovak/glide-plugin-node/commit/32ece41f82224f65e17a1b245f24d29fa7d6a6cd))

## [2.2.0](https://github.com/ivannovak/glide-plugin-node/compare/v2.1.0...v2.2.0) (2025-11-24)


### Features

* migrate to Glide v2 module path ([564b2ee](https://github.com/ivannovak/glide-plugin-node/commit/564b2ee5a062a65d7753e10ab3502b606f4184bd))

## [2.1.0](https://github.com/ivannovak/glide-plugin-node/compare/v2.0.0...v2.1.0) (2025-11-24)


### Features

* add release workflow for cross-platform binaries ([e94c38d](https://github.com/ivannovak/glide-plugin-node/commit/e94c38d8fdf6074a17c8ec853718166ce8bfdec5))

## [2.0.0](https://github.com/ivannovak/glide-plugin-node/compare/v1.0.0...v2.0.0) (2025-11-24)


### ⚠ BREAKING CHANGES

* Plugin now uses gRPC instead of library architecture

### Code Refactoring

* migrate to gRPC architecture and cleanup legacy code ([3616c30](https://github.com/ivannovak/glide-plugin-node/commit/3616c30a23a543c5d838ef38dbbc6817c68c33e4))

## 1.0.0 (2025-11-21)


### Features

* initial Node.js plugin implementation ([41e4220](https://github.com/ivannovak/glide-plugin-node/commit/41e42206488f4ef497ab1fada18a171aae44dd1a))


### Bug Fixes

* update .gitignore to allow cmd/glide-plugin-node directory ([a4b6aca](https://github.com/ivannovak/glide-plugin-node/commit/a4b6aca21c240e6c6bdd279ccfd1eeebd2a24ca9))
* update package.json repository URL to glide-plugin-node ([09d552b](https://github.com/ivannovak/glide-plugin-node/commit/09d552b628294f7c3572f25e4bcda97b3db432b7))

## 1.0.0 (2025-11-21)


### Features

* **plugin:** initial Docker plugin extraction (Phase 6) ([545ae53](https://github.com/ivannovak/glide-plugin-docker/commit/545ae5308df59fbc0e446339fbafbce719b74892))
