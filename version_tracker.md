# Library Version Tracker

This document tracks the current versions in the codebase vs the latest stable versions available.

## Research Methodology

- Check official library websites/GitHub repositories for latest stable releases
- Verify CDN availability on jsDelivr and other CDNs used
- Document version gaps and update priorities

## Library Version Status

### 1. AlpineJS

- **Current in codebase:** 3.13.0, 3.12.3
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Medium

### 2. AnimatedCSS

- **Current in codebase:** 4.1.1
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Low

### 3. Bootstrap

- **Current in codebase:** 5.3.7, 5.3.3, 5.3.2, 5.3.1, 5.3.0, 5.2.3
- **Latest stable:** 5.3.8
- **Status:** NEEDS UPDATE (5.3.8 > 5.3.7)
- **Priority:** High

### 4. Bootstrap Icons

- **Current in codebase:** 1.11.3, 1.11.2, 1.11.0, 1.10.5, 1.10.2, 1.9.1
- **Latest stable:** 1.13.1
- **Status:** NEEDS UPDATE (1.13.1 > 1.11.3)
- **Priority:** High

### 5. ChartsCSS

- **Current in codebase:** 0.9.3
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Low

### 6. FontAwesome

- **Current in codebase:** 6.5.2, 6.4.2, 6.1.2
- **Latest stable:** 7.2.0
- **Status:** NEEDS UPDATE (7.2.0 > 6.5.2)
- **Priority:** High

### 7. HTMX

- **Current in codebase:** 2.0.0, 1.9.9, 1.9.6, 1.9.5, 1.9.4, 1.9.2
- **Latest stable:** 2.0.8
- **Status:** NEEDS UPDATE (2.0.8 > 2.0.0)
- **Priority:** High

### 8. jQuery

- **Current in codebase:** 3.7.1, 3.6.4
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Medium

### 9. jQuery DataTables

- **Current in codebase:** 1.13.4
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Medium

### 10. JQTree

- **Current in codebase:** 1.8.0, 1.7.0
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Low

### 11. Notify

- **Current in codebase:** 0.4.2
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Low

### 12. SweetAlert2

- **Current in codebase:** 11, 10
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Medium

### 13. TailwindCSS

- **Current in codebase:** 3.4.4, 3.3.3
- **Latest stable:** 4.2.1
- **Status:** NEEDS UPDATE (4.2.1 > 3.4.4)
- **Priority:** High

### 14. Trumbowyg

- **Current in codebase:** 2.27.3
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Medium

### 15. VueJS

- **Current in codebase:** 3
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Medium

### 16. Vue Element Plus

- **Current in codebase:** 2.3.8
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Medium

### 17. Vue Trumbowyg

- **Current in codebase:** 4.0.0
- **Latest stable:** [NEED TO RESEARCH]
- **Status:** [NEED TO RESEARCH]
- **Priority:** Low

## Update Strategy

1. Start with highest priority libraries (Bootstrap, FontAwesome, HTMX, Bootstrap Icons, TailwindCSS)
2. Research latest versions using official sources
3. Verify CDN availability
4. Add new version functions maintaining reverse alphabetical order
5. Add comprehensive tests for each new version
6. Run full test suite after each library update
7. Update README.md documentation

## Notes

- All libraries use jsDelivr CDN primarily
- Some libraries use other CDNs (e.g., FontAwesome uses @fortawesome)
- Maintain reverse alphabetical order for version functions
- Each new version needs its own test function
- Tests should verify URL contains expected version string and CDN prefix
