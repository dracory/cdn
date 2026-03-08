# System Patterns

## Architecture Overview

This is a simple Go package that follows a functional programming pattern for CDN URL management. The system is designed to be lightweight, dependency-free, and easily extensible.

## Key Patterns

### 1. Functional API Design

- Each CDN resource is exposed as a simple function
- Functions return string URLs
- No state management or complex data structures

### 2. Naming Convention

- Format: `LibraryName_Version()` (e.g., `Jquery_3_7_1()`)
- CSS variants: `LibraryName_Version_CSS()` (e.g., `Notiflix_3_2_8_CSS()`)
- Reverse alphabetical ordering in files for latest versions on top

### 3. Environment Variable Support

- `CDN_URL_PREFIX` environment variable allows custom CDN base URLs
- `cdnBase()` function handles the logic for custom vs default URLs
- Enables easy switching between CDNs (jsDelivr, unpkg, etc.)

### 4. Test-Driven Development

- Every CDN function has a corresponding test file
- Tests verify URL contains expected version strings
- All tests must pass before deployment

### 5. File Organization

- Each library gets its own `.go` and `.test.go` file pair
- Files are kept small and focused
- Easy to add new libraries by following the pattern

## Component Relationships

```
cdn.go (base function)
    ↓
Library files (jquery.go, bootstrap.go, etc.)
    ↓
Test files (jquery_test.go, bootstrap_test.go, etc.)
    ↓
go test (verification)
```

## Extension Points

- Adding new libraries: Create new `.go` and `.test.go` files
- Version updates: Update function names and URLs
- Custom CDNs: Set `CDN_URL_PREFIX` environment variable
- New patterns: Follow existing naming and structure conventions

## Error Handling

- Minimal error handling (functions always return valid URLs)
- Tests catch invalid URLs during development
- Environment variable fallback ensures functionality
