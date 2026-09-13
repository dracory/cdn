package cdn

// Keep in reverse alphabetical order (latest version on top)

// Marked_18_0_11 returns the CDN URL for marked 18.0.11 (UMD build).
//
// marked is a low-level compiler for parsing markdown without caching or
// blocking for long periods. This is the UMD bundle that exposes a global
// `marked` object when loaded via a <script> tag.
func Marked_18_0_11() string {
	return cdnBase("https://cdn.jsdelivr.net/npm/") + "marked@18.0.11/lib/marked.umd.js"
}
