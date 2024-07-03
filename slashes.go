package slashes

import (
	"regexp"
	"strings"
)

// regex to test whether the last character is a '/'
var hasTrailingSlash = regexp.MustCompile("/$")

// regex to test whether the first character is a '/'
var hasLeadingSlash = regexp.MustCompile("^/")

// RemoveTrailingSlash removes trailing slash, if any
func RemoveTrailingSlash(path string) string {
	return strings.TrimRight(path, "/")
}

// RemoveLeadingSlash removes leading slash, if any
func RemoveLeadingSlash(path string) string {
	return strings.TrimLeft(path, "/")
}

// EnsureTrailingSlash is like AddTrailingSlash but will only ever use / since it's use for web uri's, never a Windows OS path.
func EnsureTrailingSlash(dir string) string {
	if hasTrailingSlash.MatchString(dir) {
		return dir
	}
	return dir + "/"
}

// EnsureLeadingSlash is like EnsureTrailingSlash except that it adds the leading slash if needed.
func EnsureLeadingSlash(dir string) string {
	if hasLeadingSlash.MatchString(dir) {
		return dir
	}
	return "/" + dir
}
