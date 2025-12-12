package main

import (
	"fmt"
	"os"
)

func main() {
	// need at least 1 arg (text), max 2 args (text + banner)
	if len(os.Args) < 2 || len(os.Args) > 3 {
		return // exit silently if wrong number of args
	}

	// first arg is always the text to convert
	input := os.Args[1]

	// default to standard banner
	bannerName := "standard"
	if len(os.Args) == 3 {
		// second arg is banner choice
		bannerName = os.Args[2]
	}

	// figure out which banner file to use
	var bannerPath string
	switch bannerName {
	case "standard":
		bannerPath = "banners/standard.txt"
	case "shadow":
		bannerPath = "banners/shadow.txt"
	case "thinkertoy":
		bannerPath = "banners/thinkertoy.txt"
	default:
		// invalid banner name, just exit
		return
	}

	// load the banner file
	banner, err := LoadBanner(bannerPath)
	if err != nil {
		// couldn't load banner, exit silently
		return
	}

	// convert text to ASCII art and print it
	output := RenderInput(input, banner)
	fmt.Print(output)
}


