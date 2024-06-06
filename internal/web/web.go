package web

import (
	"html"
	"strings"
)

// Scrub sanitizes user input to prevent XSS and other potential attacks.
func Scrub(input string) string {
	input = strings.TrimSpace(input)
	input = html.EscapeString(input)
	return input
}
