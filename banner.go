package main

import (
	"fmt"
	"os"
	"strings"
)

// Banner maps a rune to its ASCII-art glyph representation.
// Each glyph consists of charHeight number of lines.
type Banner map[rune][]string

const (
	firstChar  = 32  // ASCII code for ' '
	lastChar   = 126 // ASCII code for '~'
	charHeight = 8   // Number of lines per character in the banner file
)

// LoadBanner reads a banner font file (standard/shadow/thinkertoy)
// and loads its glyphs into a Banner map.
//
// The expected structure of the banner file is:
//   - For every printable ASCII character (32–126):
//     1 empty/unused line
//     followed by 8 lines representing the character's ASCII-art glyph.
//   - Therefore, each character occupies 9 lines in total.
//
// This function sequentially extracts each 9-line block and stores
// the 8 glyph lines in the Banner map.
func LoadBanner(path string) (Banner, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Split file into individual lines
	lines := strings.Split(string(data), "\n")

	banner := make(Banner)

	const blockSize = charHeight + 1 // Total lines per char block (1 empty + 8 glyph lines)

	for code := firstChar; code <= lastChar; code++ {
		blockIndex := int(code - firstChar)
		start := blockIndex * blockSize

		// Ensure there are enough lines remaining for the glyph
		if start+1+charHeight > len(lines) {
			return nil, fmt.Errorf("invalid banner file: not enough lines for char %q", rune(code))
		}

		// Skip the first line in the block (empty line), then take the next 8 glyph lines
		glyphLines := lines[start+1 : start+1+charHeight]
		banner[rune(code)] = glyphLines
	}

	return banner, nil
}