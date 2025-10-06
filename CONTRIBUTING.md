# Contributing to LeakFinder

Thank you for your interest in contributing to LeakFinder! This document provides guidelines and information for contributors.

## 🚀 Quick Start

1. **Fork the repository**
2. **Clone your fork**: `git clone https://github.com/ashutoshmishraji/leakfinder.git`
3. **Create a branch**: `git checkout -b feature/your-feature-name`
4. **Make your changes**
5. **Test your changes**: `go test ./...`
6. **Submit a pull request**

## 📋 Development Setup

### Prerequisites

- Go 1.21 or later
- Git

### Local Development

```bash
# Clone the repository
git clone https://github.com/ashutoshmishraji/leakfinder.git
cd leakfinder

# Install dependencies
go mod download

# Build the project
go build -o leakfinder leakfinder.go

# Run tests
go test ./...

# Test the binary
./leakfinder help
```

## 🎯 How to Contribute

### Reporting Bugs

1. **Check existing issues** to avoid duplicates
2. **Use the bug report template**
3. **Include**:
   - Go version (`go version`)
   - Operating system
   - LeakFinder version
   - Steps to reproduce
   - Expected vs actual behavior
   - Relevant logs/output

### Suggesting Features

1. **Check existing feature requests**
2. **Open an issue** with:
   - Clear description of the feature
   - Use case and motivation
   - Possible implementation approach
   - Examples of usage

### Code Contributions

#### Areas for Contribution

- **New leak detection patterns**
- **Performance optimizations**
- **Additional output formats**
- **Integration with monitoring systems**
- **Documentation improvements**
- **Test coverage improvements**

#### Code Style Guidelines

**Go Code:**
- Follow `gofmt` formatting
- Use meaningful variable names
- Add comments for exported functions
- Keep functions under 50 lines when possible
- Use proper error handling
- Follow Go best practices

**Example:**
```go
// AnalyzeGoroutines performs comprehensive goroutine leak analysis
func (ldt *LeakDetectionToolkit) AnalyzeGoroutines() (*Analysis, error) {
    if ldt.config == nil {
        return nil, fmt.Errorf("configuration not initialized")
    }
    
    // Implementation...
}
```

#### Testing Guidelines

- **Write tests** for new features
- **Update existing tests** when modifying functionality
- **Include edge cases** in test scenarios
- **Use table-driven tests** for multiple scenarios

```go
func TestAnalyzeGoroutines(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected int
        wantErr  bool
    }{
        {"normal case", "goroutine 1 [running]:", 1, false},
        {"empty input", "", 0, false},
        // Add more test cases...
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation...
        })
    }
}
```

## 📝 Pull Request Process

1. **Update documentation** if needed
2. **Add/update tests** for your changes
3. **Ensure all tests pass**: `go test ./...`
4. **Run linting**: `go vet ./...`
5. **Update CHANGELOG.md** if applicable
6. **Write clear commit messages**

### Commit Message Format

```
type(scope): brief description

Detailed explanation if needed

Fixes #123
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes
- `refactor`: Code refactoring
- `test`: Test additions/modifications
- `chore`: Maintenance tasks

**Examples:**
```
feat(analysis): add memory pattern detection
fix(gc): correct GC effectiveness calculation
docs(readme): update installation instructions
```

## 🧪 Testing

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detection
go test -race ./...

# Run specific test
go test -run TestAnalyzeGoroutines
```

### Test Categories

- **Unit tests**: Test individual functions
- **Integration tests**: Test component interactions
- **End-to-end tests**: Test complete workflows

## 📚 Documentation

### Code Documentation

- Use **GoDoc** format for function comments
- Include **examples** in documentation
- Document **configuration options**
- Explain **complex algorithms**

### README Updates

When adding features, update:
- Usage examples
- Command descriptions
- Configuration options
- Performance notes

## 🔍 Code Review Guidelines

### For Contributors

- **Keep PRs focused** on a single feature/fix
- **Write descriptive PR descriptions**
- **Respond to feedback** promptly
- **Update your PR** based on review comments

### For Reviewers

- **Be constructive** and helpful
- **Focus on code quality** and maintainability
- **Check for test coverage**
- **Verify documentation updates**

## 🏷️ Release Process

1. **Update version** in relevant files
2. **Update CHANGELOG.md**
3. **Create release PR**
4. **Tag release** after merge
5. **GitHub Actions** will build and publish binaries

## 🤝 Community Guidelines

- **Be respectful** and inclusive
- **Help newcomers** get started
- **Share knowledge** and best practices
- **Follow the Code of Conduct**

## 📞 Getting Help

- **GitHub Issues**: For bugs and feature requests
- **GitHub Discussions**: For questions and general discussion
- **Documentation**: Check README and code comments

## 🎉 Recognition

Contributors will be:
- **Listed in CONTRIBUTORS.md**
- **Mentioned in release notes**
- **Credited in documentation**

Thank you for contributing to LeakFinder! 🚀
