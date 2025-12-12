package main

import (
	"fmt"
	"strings"
)

// validateInput checks if all characters in input are printable ASCII
// only allows characters from space (32) to ~ (126), plus newlines
func validateInput(input string) error {
	// keep track of bad characters
	unsupported := make([]rune, 0, len(input))

	// check each character in the input
	for _, r := range input {
		// newlines are always ok
		if r == '\n' {
			continue
		}

		// check if character is in our supported range
		if r < firstChar || r > lastChar {
			unsupported = append(unsupported, r)
		}
	}

	// if no bad characters, we're good
	if len(unsupported) == 0 {
		return nil
	}

	// build error message with the bad characters
	out := make([]string, 0, len(unsupported))
	for _, r := range unsupported {
		out = append(out, fmt.Sprintf("%q", r))
	}

	// return appropriate error message
	if len(unsupported) == 1 {
		return fmt.Errorf("unsupported character: %s", out[0])
	}

	return fmt.Errorf("unsupported characters: %s", strings.Join(out, ", "))
}