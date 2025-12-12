package main

import "strings"

// RenderLine takes a string and makes ASCII art from it
// builds it row by row (each character is 8 rows tall)
func RenderLine(s string, b Banner) string {
	const height = charHeight

	var builder strings.Builder

	// make empty glyph for characters we don't have
	empty := make([]string, height)

	// go through each row (0 to 7)
	for row := 0; row < height; row++ {
		// for each character in the input string
		for _, ch := range s {
			glyph, ok := b[ch]
			if !ok {
				// character not in banner, use empty space
				glyph = empty
			}
			// add this character's row to the line
			builder.WriteString(glyph[row])
		}
		// end of row, add newline
		builder.WriteRune('\n')
	}

	return builder.String()
}

// decodeEscapedNewlines converts \n to actual newlines
// when user types "Hello\nWorld" we want real newlines
func decodeEscapedNewlines(s string) string {
	runes := []rune(s)
	out := make([]rune, 0, len(runes))

	// go through each character
	for i := 0; i < len(runes); i++ {
		// if we see \ followed by n, make it a real newline
		if runes[i] == '\\' && i+1 < len(runes) && runes[i+1] == 'n' {
			out = append(out, '\n')
			i++ // skip the n since we already used it
		} else {
			// normal character, just copy it
			out = append(out, runes[i])
		}
	}

	return string(out)
}

// RenderInput is the main function that handles everything
// takes user input and converts it to ASCII art
func RenderInput(input string, b Banner) string {
	// first convert \n strings to real newlines
	input = decodeEscapedNewlines(input)

	// if input is empty, return empty
	if input == "" {
		return ""
	}

	// if no newlines, just render as single line
	if !strings.Contains(input, "\n") {
		return RenderLine(input, b)
	}

	var builder strings.Builder

	// split input by newlines to handle multiple lines
	parts := strings.Split(input, "\n")
	hadText := false

	// process each line
	for i, part := range parts {
		last := i == len(parts)-1

		if part != "" {
			// non-empty line, render it
			asciiBlock := RenderLine(part, b)
			builder.WriteString(asciiBlock)
			hadText = true

		} else {
			// empty line handling
			if !last {
				// empty line in middle, add blank line
				builder.WriteRune('\n')
			} else if hadText {
				// empty line at end, keep the newline
				builder.WriteRune('\n')
			}
		}
	}

	return builder.String()
}