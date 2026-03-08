# Agent Instructions

This file contains instructions for AI agents working on the cdn package.

## Overview

The cdn package provides CDN URLs for popular frontend libraries. The goal is to keep library versions up to date while maintaining backward compatibility and comprehensive testing.

## Current Status

### ✅ Completed Updates (March 8, 2026)

- **Bootstrap**: 5.3.7 → 5.3.8
- **FontAwesome**: 6.5.2 → 7.2.0
- **HTMX**: 2.0.0 → 2.0.8
- **Bootstrap Icons**: 1.11.3 → 1.13.1 (already had latest)
- **TailwindCSS**: 3.4.4 → 4.2.1

### 🔄 Remaining Libraries to Update

- AlpineJS (3.13.0)
- AnimatedCSS (4.1.1)
- jQuery (3.7.1)
- jQuery DataTables (1.13.4)
- JQTree (1.8.0)
- Notify (0.4.2)
- SweetAlert2 (11)
- Trumbowyg (2.27.3)
- VueJS (3)
- Vue Element Plus (2.3.8)
- Vue Trumbowyg (4.0.0)
- ChartsCSS (0.9.3)

## Update Process

### 1. Research Latest Versions

Use the jsDelivr API to check for latest versions:

```powershell
Invoke-WebRequest -Uri 'https://data.jsdelivr.com/v1/package/npm/[package-name]' -UseBasicParsing | Select-Object -ExpandProperty Content | ConvertFrom-Json | Select-Object -ExpandProperty tags | Format-List
```

### 2. Verify CDN Availability

Test that the latest version is available on the CDN:

```powershell
Invoke-WebRequest -Uri 'https://cdn.jsdelivr.net/npm/[package-name]@[version]/[file]'
```

### 3. Add New Version Functions

- Add functions in reverse alphabetical order (latest version on top)
- Follow existing naming conventions: `LibraryName_Version_Major_Minor_Patch()`
- Use appropriate CDN URLs (jsDelivr, unpkg, etc.)

### 4. Add Comprehensive Tests

- Create test functions for each new version
- Verify URL contains expected version string
- Test CDN prefix functionality
- Follow existing test patterns

### 5. Update Documentation

- Update README.md with new version information
- Keep version lists current and accurate
- Update LIBRARY_VERSIONS.md status

### 6. Run Full Test Suite

- Execute `go test ./...` to ensure all tests pass
- Verify no regressions in existing functionality

## File Structure

### Library Files

Each library has:

- `[library_name].go` - Contains version functions
- `[library_name]_test.go` - Contains corresponding tests

### Documentation Files

- `README.md` - Public documentation with latest versions
- `LIBRARY_VERSIONS.md` - Internal tracking and maintenance guide

## Best Practices

### Version Management

- Always add new versions rather than replacing existing ones
- Maintain backward compatibility
- Keep multiple recent versions available

### Testing

- Each new version must have corresponding tests
- Tests should verify both URL format and CDN prefix functionality
- Run full test suite after each library update

### Documentation

- Keep README.md current with latest versions
- Document any breaking changes or migration notes
- Update LIBRARY_VERSIONS.md regularly

## Troubleshooting

### Common Issues

1. **Duplicate function names**: Ensure unique function names for each version
2. **CDN availability**: Verify CDN URLs are accessible
3. **Test failures**: Check URL format and version strings
4. **Documentation sync**: Keep README.md in sync with code

### Resolution Steps

1. Verify the issue with manual testing
2. Check CDN availability directly
3. Review function naming and URL format
4. Update tests and documentation as needed

## Notes

- All libraries use jsDelivr CDN primarily
- Some libraries use other CDNs (e.g., FontAwesome uses @fortawesome)
- Maintain reverse alphabetical order for version functions
- Each new version needs its own test function
- Tests should verify URL contains expected version string and CDN prefix
