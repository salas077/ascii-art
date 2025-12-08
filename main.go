package main

import (
	"fmt"
	"os"
)

func main() {
	// 1. If no arguments are provided (only the program name), exit
	if len(os.Args) < 2 {
		return
	}

	// 2. Read the input string from command-line argument
	input := os.Args[1]

	// 3. If the input is an empty string "", exit
	if input == "" {
		return
	}

	// 4. TEMPORARY: Just print the input so we know the program works.
	//    This will be replaced later with ASCII-art logic.
	fmt.Println("Input:", input)

}
