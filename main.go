package main

import (
	"fmt"
	"os"

	asciiArt "ascii-art/ascii-funcs" // Importing custom package for ASCII art functions
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}

	if len(os.Args) == 3 {
		text := os.Args[1]
		banner := os.Args[2]

		asciiChars, err := asciiArt.LoadAsciiChars(banner + ".txt") // Load ASCII characters from the specified file
		if err != nil {
			fmt.Println("Error!: Did you mean 'thinkertoy or 'shadow' or 'standard?'", err) // Print error message and suggestion
			return                                                                          // Exit the program if there's an error
		}
		asciiArt.ProcessArguments(text, asciiChars) // Process the first command-line argument with loaded ASCII characters
	}
}
