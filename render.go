package main

import "strings"

// RenderLine generates the ASCII-art representation of a *single*
// logical line of text (a string that does NOT contain real '\n').
// For each of the 8 rows of the glyphs, the function concatenates
// the corresponding row of each character.
func RenderLine(s string, b Banner) string {
	const height = charHeight

	var builder strings.Builder

	// Create empty glyph once for unknown characters
	empty := make([]string, height)

	for row := 0; row < height; row++ {
		for _, ch := range s {
			glyph, ok := b[ch]
			if !ok {
				// If character not found, use empty glyph
				glyph = empty
			}
			builder.WriteString(glyph[row])
		}
		builder.WriteRune('\n')
	}

	return builder.String()
}

// decodeEscapedNewlines converts escape sequences "\n"
// (two literal characters: '\' and 'n') into actual newline runes.
func decodeEscapedNewlines(s string) string {
	runes := []rune(s)
	out := make([]rune, 0, len(runes))

	for i := 0; i < len(runes); i++ {
		if runes[i] == '\\' && i+1 < len(runes) && runes[i+1] == 'n' {
			out = append(out, '\n')
			i++ // Skip the 'n'
		} else {
			out = append(out, runes[i])
		}
	}

	return string(out)
}

// RenderInput processes the entire user input (which may include
// the user-typed literal sequence "\n") and produces the final,
// multi-line ASCII-art output according to the project examples.
//
// Behavior summary:
//  1. Converts all "\n" sequences into real newline characters.
//  2. If the result contains no newline at all → render as a single line.
//  3. If newlines exist → split into logical lines and render each one.
//  4. Empty logical lines ("") produce a single blank ASCII-art line.
func RenderInput(input string, b Banner) string {
	// Step 1: convert "\n" into real newline runes.
	input = decodeEscapedNewlines(input)

	// Step 2: if input is completely empty, return an empty output.
	if input == "" {
		return ""
	}

	// If there are no real '\n' characters, this is a single-line case.
	if !strings.Contains(input, "\n") {
		return RenderLine(input, b)
	}

	var builder strings.Builder

	// Step 3: split into logical lines.
	parts := strings.Split(input, "\n")
	hadText := false

	for i, part := range parts {
		last := i == len(parts)-1

		if part != "" {
			// Non-empty line → render it normally.
			asciiBlock := RenderLine(part, b)
			builder.WriteString(asciiBlock)
			hadText = true

		} else {
			// Empty logical line.
			if !last {
				// Empty line between text lines → output one blank line.
				builder.WriteRune('\n')
			} else if hadText {
				// Last line was empty but previous had content → maintain final newline.
				builder.WriteRune('\n')
			}
		}
	}

	return builder.String()
}