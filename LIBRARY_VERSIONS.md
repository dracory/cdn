# Library Version Management

This document serves as both a version tracker and maintenance guide for the cdn package.

## Current Status (Updated: March 8, 2026)

### ✅ Completed Updates

- **Bootstrap**: 5.3.7 → 5.3.8 ✅
- **FontAwesome**: 6.5.2 → 7.2.0 ✅
- **HTMX**: 2.0.0 → 2.0.8 ✅
- **Bootstrap Icons**: 1.11.3 → 1.13.1 ✅ (already had latest)
- **TailwindCSS**: 3.4.4 → 4.2.1 ✅

### 🔄 Remaining Libraries to Update

| Library           | Current | Latest   | Status  | Priority |
| ----------------- | ------- | -------- | ------- | -------- |
| AlpineJS          | 3.13.0  | 3.15.8   | ✅ Done | Medium   |
| AnimatedCSS       | 4.1.1   | 4.1.1    | ✅ Done | Low      |
| jQuery            | 3.7.1   | 4.0.0    | ✅ Done | Medium   |
| jQuery DataTables | 1.13.4  | 2.3.7    | ✅ Done | Medium   |
| JQTree            | 1.8.0   | 1.8.3    | ✅ Done | Low      |
| Notify            | 0.4.2   | 3.0.0    | ✅ Done | Low      |
| Slazy             | -       | 0.5.0    | ✅ Done | Low      |
| SweetAlert2       | 11      | 11.26.22 | ✅ Done | Medium   |
| Trumbowyg         | 2.27.3  | 2.31.0   | ✅ Done | Medium   |
| VueJS             | 3       | 3.5.30   | ✅ Done | Medium   |
| Vue Element Plus  | 2.3.8   | 2.13.5   | ✅ Done | Medium   |
| Vue Trumbowyg     | 4.0.0   | 4.0.0    | ✅ Done | Low      |
| ChartsCSS         | 0.9.3   | 1.2.0    | ✅ Done | Low      |

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

### 6. Run Full Test Suite

- Execute `go test ./...` to ensure all tests pass
- Verify no regressions in existing functionality

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
- Update this version tracker regularly

## Automation Opportunities

Consider automating the following tasks:

1. **Version checking**: Script to check for latest versions of all libraries
2. **CDN availability testing**: Automated verification of CDN URLs
3. **Test generation**: Template-based test creation for new versions
4. **Documentation updates**: Automated README.md updates

## Monitoring

### Regular Checks

- Monthly review of library versions
- Monitor library release schedules
- Watch for security updates and patches

### Alerts

- Set up notifications for major version releases
- Monitor CDN availability and performance
- Track usage statistics if available

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
