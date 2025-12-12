package main

import (
	"fmt"
	"os"
)

func main() {
	// Allowed argument scenarios:
	//  - 1 arg  -> message, banner = standard
	//  - 2 args -> message, banner = standard/shadow/thinkertoy
	if len(os.Args) < 2 || len(os.Args) > 3 {
		return
	}

	input := os.Args[1]

	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	// Choose banner file based on the bannerName argument.
	var bannerPath string
	switch bannerName {
	case "standard":
		bannerPath = "banners/standard.txt"
	case "shadow":
		bannerPath = "banners/shadow.txt"
	case "thinkertoy":
		bannerPath = "banners/thinkertoy.txt"
	default:
		// If an unknown banner name is provided, do nothing and exit.
		return
	}

	banner, err := LoadBanner(bannerPath)
	if err != nil {
		return
	}

	output := RenderInput(input, banner)
	fmt.Print(output)
}


