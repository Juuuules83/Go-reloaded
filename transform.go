package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidText = errors.New("texte invalide")

var directivePattern = regexp.MustCompile(`\((hex|bin|up|low|cap)\s*,\s*(\d+)\)`)
var directiveWordPattern = regexp.MustCompile(`^\((hex|bin|up|low|cap)(?:,(\d+))?\)$`)

func ProcessText(input string) (string, error) {
	input = directivePattern.ReplaceAllString(input, "($1,$2)")
	words := strings.Fields(input)
	result := []string{}

	for _, word := range words {
		match := directiveWordPattern.FindStringSubmatch(word)
		if match == nil {
			result = append(result, word)
			continue
		}

		count := 1
		if match[2] != "" {
			number, err := strconv.Atoi(match[2])
			if err != nil || number < 1 {
				return "", fmt.Errorf("%w : nombre incorrect dans %q", ErrInvalidText, word)
			}
			count = number
		}

		switch match[1] {
		case "hex":
			if err := convertNumber(&result, 16); err != nil {
				return "", err
			}
		case "bin":
			if err := convertNumber(&result, 2); err != nil {
				return "", err
			}
		case "up", "low", "cap":
			if err := changeCase(result, match[1], count); err != nil {
				return "", err
			}
		}
	}

	text := strings.Join(result, " ")
	text = fixApostrophes(text)
	text = fixPunctuation(text)
	text = fixArticles(text)
	return text, nil
}
