package main

import (
	"fmt"
	"strings"
	"unicode"
)

func changeCase(words []string, kind string, count int) error {
	if len(words) < count {
		return fmt.Errorf("%w : pas assez de mots avant la directive", ErrInvalidText)
	}

	for i := len(words) - count; i < len(words); i++ {
		switch kind {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = capitalize(words[i])
		}
	}
	return nil
}

func capitalize(word string) string {
	runes := []rune(strings.ToLower(word))
	if len(runes) == 0 {
		return word
	}
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
