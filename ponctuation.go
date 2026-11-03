package main

import (
	"regexp"
	"strings"
)

func fixPunctuation(text string) string {
	pattern := regexp.MustCompile(`\s*([.,!?;:]+)\s*`)
	text = pattern.ReplaceAllString(text, "$1 ")
	return strings.TrimSpace(text)
}
