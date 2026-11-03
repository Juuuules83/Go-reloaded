package main

import (
	"fmt"
	"strconv"
)

func convertNumber(words *[]string, base int) error {
	if len(*words) == 0 {
		return fmt.Errorf("%w : aucun nombre avant la directive", ErrInvalidText)
	}

	index := len(*words) - 1
	number, err := strconv.ParseInt((*words)[index], base, 64)
	if err != nil {
		return fmt.Errorf("%w : conversion du nombre %q", ErrInvalidText, (*words)[index])
	}

	(*words)[index] = strconv.FormatInt(number, 10)
	return nil
}
