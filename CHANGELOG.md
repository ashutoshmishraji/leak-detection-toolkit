# Changelog

All notable changes to LeakFinder will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Initial release of LeakFinder
- Goroutine leak detection with advanced pattern recognition
- Memory leak analysis with heap tracking
- GC analysis to distinguish real leaks from GC delays
- Production optimizations (HTTP pooling, string interning, pre-compiled regexes)
- Focus options for goroutines-only or memory-only analysis
- Multiple detection algorithms (Stack patterns, Uber LeakProf)
- Real-time monitoring with configurable intervals
- Datadog-inspired memory pattern analysis
- Comprehensive reporting with detailed recommendations
- Custom pattern configuration via flags and environment variables
- GitHub Actions CI/CD pipeline
- Cross-platform binary builds
- Comprehensive documentation and examples

### Features
- **Commands**: `check`, `monitor`, `analyze`, `compare`, `filter`, `patterns`, `channels`
- **GC Analysis**: `-force-gc` flag with customizable parameters
- **Focus Options**: `-goroutines-only`, `-memory-only` flags
- **Pattern Configuration**: Custom suspicious and safe patterns
- **Performance Optimized**: HTTP connection pooling, pre-compiled regexes
- **Reporting**: Automatic report generation in `./reports/` directory
- **Shell Script**: Enhanced runner with advanced options
- **Cross-Platform**: Linux, macOS, Windows support

### Technical Details
- Go 1.21+ support
- Production-ready performance optimizations
- Comprehensive test coverage
- Static analysis with staticcheck
- Multi-platform CI/CD builds
- MIT License

## [1.0.0] - 2025-01-06

### Added
- Initial public release
- Core leak detection functionality
- Documentation and examples
- CI/CD pipeline
- Cross-platform support
