package cdn

// Keep in reverse alphabetical order (latest version on top)

// DOMPurify_3_4_14 returns the CDN URL for DOMPurify 3.4.14 (minified build).
//
// DOMPurify is a DOM-only, super-fast, uber-tolerant XSS sanitizer for HTML,
// MathML and SVG. This is the minified UMD bundle that exposes a global
// `DOMPurify` object when loaded via a <script> tag.
func DOMPurify_3_4_14() string {
	return cdnBase("https://cdn.jsdelivr.net/npm/") + "dompurify@3.4.14/dist/purify.min.js"
}
