package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	charHeight = 8   // each ASCII-art character is 8 lines tall
	firstChar  = 32  // ASCII 32 is ' '
	lastChar   = 126 // ASCII 126 is '~'
)

func main() {
	// 1. Check if we have at least one argument (the text)
	if len(os.Args) < 2 {
		return
	}

	// 2. Read the input text
	input := os.Args[1]
	if input == "" {
		return
	}

	// 3. Choose which banner file to use
	bannerFile := "banners/standard.txt" // default banner

	// If there is a second argument, use it as banner name
	if len(os.Args) >= 3 {
		bannerName := os.Args[2]

		switch bannerName {
		case "standard":
			bannerFile = "banners/standard.txt"
		case "shadow":
			bannerFile = "banners/shadow.txt"
		case "thinkertoy":
			bannerFile = "banners/thinkertoy.txt"
		default:
			fmt.Println("Unknown banner, using standard.")
		}
	}

	// 4. Load the font from the selected banner file
	font, err := loadFont(bannerFile)
	if err != nil {
		log.Fatal(err)
	}

	// 5. Convert literal "\n" into real newline
	input = strings.ReplaceAll(input, `\n`, "\n")

	// 6. Split the input into lines
	lines := strings.Split(input, "\n")

	// 7. For each line
	for _, line := range lines {
		// If line is empty → print blank ASCII block
		if line == "" {
			printEmptyBlock()
			continue
		}

		// Otherwise print the line as ASCII-art
		printWordAsAscii(line, font)
	}
}

// loadFont reads the banner file and builds a map:
// key   = ASCII code (byte) of the character
// value = slice of 8 strings representing that character in ASCII-art
func loadFont(path string) (map[byte][]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Split the whole file into lines
	lines := strings.Split(string(data), "\n")

	font := make(map[byte][]string)

	// Number of characters in the banner (from ASCII 32 to 126)
	totalChars := int(lastChar - firstChar + 1) // 95 characters

	// For each character
	for i := 0; i < totalChars; i++ {
		ascii := byte(firstChar + i)

		// Each character block has 8 lines + 1 empty line separator
		start := i * (charHeight + 1)

		var art []string
		for j := 0; j < charHeight; j++ {
			art = append(art, lines[start+j])
		}

		font[ascii] = art
	}

	return font, nil
}

// printWordAsAscii prints one word as ASCII-art using the given font.
func printWordAsAscii(word string, font map[byte][]string) {
	// for each row (0 to 7)
	for row := 0; row < charHeight; row++ {

		line := ""

		// for each character in the input word
		for i := 0; i < len(word); i++ {
			ch := word[i]

			// ignore characters outside the banner range
			if ch < firstChar || ch > lastChar {
				continue
			}

			art := font[ch] // 8-line ASCII art block for this character

			line += art[row]
		}

		fmt.Println(line)
	}
}

// printEmptyBlock prints an empty ASCII-art line (8 empty lines)
func printEmptyBlock() {
	for i := 0; i < charHeight; i++ {
		fmt.Println()
	}
}
