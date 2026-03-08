# Project Brief

## Overview

This is a Go package that provides CDN URLs for popular JavaScript libraries, CSS frameworks, and other web resources. The package allows developers to easily reference CDN-hosted assets in their Go applications.

## Purpose

- Centralize CDN URL management for web development projects
- Provide version-specific CDN URLs for popular libraries
- Support custom CDN URL prefixes through environment variables
- Include comprehensive test coverage for all CDN functions

## Scope

The package includes CDN URLs for:

- JavaScript frameworks (Vue.js, jQuery, HTMX, etc.)
- CSS frameworks (Bootstrap, Tailwind CSS, FontAwesome, etc.)
- UI libraries and components
- Utility libraries and plugins

## Key Requirements

- Each CDN resource must have a corresponding test file
- Functions should follow naming convention: `LibraryName_Version()`
- URLs should point to specific versions for stability
- Support for both JS and CSS files where applicable
- Environment variable support for custom CDN prefixes
