package main

import "regexp"

func fixArticles(text string) string {
	pattern := regexp.MustCompile(`(?i)\ba\s+([aeiouh])`)
	return pattern.ReplaceAllStringFunc(text, func(match string) string {
		if match[0] == 'A' {
			return "An" + match[1:]
		}
		return "an" + match[1:]
	})
}
