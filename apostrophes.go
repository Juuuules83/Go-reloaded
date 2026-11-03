package main

import "regexp"

func fixApostrophes(text string) string {
	pattern := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return pattern.ReplaceAllString(text, "'$1'")
}
