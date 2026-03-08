# Technical Context

## Technologies Used

- **Go 1.21+** - Primary programming language
- **Standard Library Only** - No external dependencies
- **Go Testing Framework** - Built-in testing with `testing` package
- **strings package** - For URL validation in tests

## Development Setup

- **Go Modules** - Managed via `go.mod`
- **Standard Go Toolchain** - `go test`, `go build`, `go mod tidy`
- **No Build Tools Required** - Simple package structure

## Technical Constraints

- **No External Dependencies** - Pure Go implementation
- **String-Based URLs** - Functions return plain strings
- **Environment Variable Support** - Custom CDN prefixes via `CDN_URL_PREFIX`
- **Version-Specific URLs** - Hardcoded to specific library versions

## Dependencies

```
go.mod
module github.com/dracory/cdn

go 1.21
```

## Testing Strategy

- **Unit Tests** - Each function has corresponding test
- **String Validation** - Tests verify URL contains expected version strings
- **No Network Calls** - Tests are purely functional
- **100% Coverage** - All functions must have tests

## Build and Deployment

- **Simple Package** - No complex build process
- **Go Module** - Published as standard Go module
- **Cross-Platform** - Works on any platform Go supports
- **No Runtime Dependencies** - Self-contained package

## Performance Characteristics

- **Zero Runtime Overhead** - Functions return static strings
- **Environment Variable Check** - Single `os.Getenv()` call per function
- **Memory Efficient** - No caching or state management
- **Fast Execution** - Direct string returns

## Code Quality Standards

- **Go Formatting** - Standard `gofmt` formatting
- **Naming Conventions** - PascalCase for functions, camelCase for variables
- **Documentation** - Inline comments for version ordering
- **Test Coverage** - Comprehensive test suite required
