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
	// Check if we have at least one argument
	if len(os.Args) < 2 {
		return
	}

	// Read the input (we won't use it yet, we just test 'A')
	input := os.Args[1]
	if input == "" {
		return
	}

	// Load the font from the banner file
	font, err := loadFont("standard.txt")
	if err != nil {
		log.Fatal(err)
	}

	// TEST: print the ASCII-art for the character 'A'
	art := font['A']

	for _, line := range art {
		fmt.Println(line)
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
