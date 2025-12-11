package main

import (
	"testing"
)

func TestLoadFontStandard(t *testing.T) {
	// Call loadFont with the standard banner file
	font, err := loadFont("standard.txt")
	if err != nil {
		t.Fatalf("loadFont returned an error: %v", err)
	}

	// Check that the map is not empty
	if len(font) == 0 {
		t.Fatalf("font map is empty")
	}

	// Check that we have the expected number of characters
	expectedChars := int(lastChar - firstChar + 1)
	if len(font) != expectedChars {
		t.Errorf("expected %d characters, got %d", expectedChars, len(font))
	}

	// For each character in the range, check that:
	// - it exists in the map
	// - it has exactly charHeight lines
	for c := firstChar; c <= lastChar; c++ {
		art, ok := font[byte(c)]
		if !ok {
			t.Errorf("missing character %q (%d) in font map", c, c)
			continue
		}

		if len(art) != charHeight {
			t.Errorf("character %q has %d lines, expected %d", c, len(art), charHeight)
		}
	}
}
