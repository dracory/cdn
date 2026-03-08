# Product Context

## What This Project Solves

This project solves the problem of managing CDN URLs in Go web applications. Instead of hardcoding CDN URLs throughout your codebase, this package provides a centralized, version-controlled way to reference popular web libraries and frameworks.

## Current State

The project successfully provides CDN URLs for 20+ popular libraries including:

- JavaScript frameworks (Vue.js, jQuery, HTMX)
- CSS frameworks (Bootstrap, Tailwind CSS, FontAwesome)
- UI components (SweetAlert2, DataTables, Trumbowyg)
- Utility libraries (Notiflix, Notify)

## Recent Addition: Notiflix

**Date Added:** March 8, 2026
**Version:** 3.2.8
**Files Added:**

- `notiflix.go` - Contains CDN functions for JavaScript and CSS
- `notiflix_test.go` - Comprehensive test coverage

**Functions Added:**

- `Notiflix_3_2_8()` - Returns JavaScript URL: `https://cdn.jsdelivr.net/npm/notiflix@3.2.8/dist/notiflix-aio-3.2.8.min.js`
- `Notiflix_3_2_8_CSS()` - Returns CSS URL: `https://cdn.jsdelivr.net/npm/notiflix@3.2.8/src/notiflix.min.css`

## User Experience Goals

- Developers can easily import and use CDN URLs in their Go templates
- Version consistency across projects
- Easy migration to new library versions
- Comprehensive test coverage ensures URL validity

## Integration Points

The package integrates with:

- Go web frameworks (Gin, Echo, etc.)
- HTML template engines
- Build systems that need CDN asset references
