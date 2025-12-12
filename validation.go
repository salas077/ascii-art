package main

import (
	"fmt"
	"strings"
)

// validateInput checks if all characters in the input are supported
// (within the ASCII range 32-126, plus newlines are allowed)
func validateInput(input string) error {
	// Pre-allocate slice for better performance
	unsupported := make([]rune, 0, len(input))

	for _, r := range input {
		if r == '\n' {
			continue
		}

		if r < firstChar || r > lastChar {
			unsupported = append(unsupported, r)
		}
	}

	if len(unsupported) == 0 {
		return nil
	}

	// Build readable list: 'Γ', 'ε', '🙂'
	out := make([]string, 0, len(unsupported))
	for _, r := range unsupported {
		out = append(out, fmt.Sprintf("%q", r))
	}

	// Singular or plural?
	if len(unsupported) == 1 {
		return fmt.Errorf("unsupported character: %s", out[0])
	}

	return fmt.Errorf("unsupported characters: %s", strings.Join(out, ", "))
}