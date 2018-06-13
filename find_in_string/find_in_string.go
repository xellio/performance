package performance

import (
	"bytes"
	"regexp"
	"strings"
)

func stringContains(haystack string, needle string) bool {
	return strings.Contains(haystack, needle)
}

func byteContains(haystack []byte, needle []byte) bool {
	return bytes.Contains(haystack, needle)
}

func byteContainsWithCast(haystack string, needle string) bool {
	return bytes.Contains([]byte(haystack), []byte(needle))
}

func stringContainsWithCast(haystack []byte, needle []byte) bool {
	return strings.Contains(string(haystack), string(needle))
}

func regexpMatchString(pattern string, haystack string) bool {
	match, _ := regexp.MatchString(pattern, haystack)
	return match
}

func regexpMatch(pattern string, haystack []byte) bool {
	match, _ := regexp.Match(pattern, haystack)
	return match
}

func regexpMatchCompiled(pattern string, haystack []byte) bool {
	r, _ := regexp.Compile(pattern)
	return r.Match(haystack)
}
