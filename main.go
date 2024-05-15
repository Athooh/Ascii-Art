package main

import (
	asciiArt "ascii-art/ascii-funcs" // Importing custom package for ASCII art functions
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 { 
		fmt.Println("Usage: go run . <string>") 
	}

	args := os.Args[1] 

	if len(os.Args) == 3 { 
		args2 := os.Args[2] 
		asciiChars, err := asciiArt.LoadAsciiChars(args2) // Load ASCII characters from the specified file
		if err != nil { 
			fmt.Println("Error!: Did you mean 'thinkertoy.txt or 'shadow.txt'?", err) // Print error message and suggestion
			return // Exit the program if there's an error
		}
		asciiArt.ProcessArguments(args, asciiChars) // Process the first command-line argument with loaded ASCII characters

	} else { // If there's not exactly two command-line arguments
		asciiChars, err := asciiArt.LoadAsciiChars("standard.txt") // Load default ASCII characters from standard.txt
		if err != nil { 
			fmt.Println("Error loading ASCII characters", err) 
			return // Exit the program if there's an error
		}
		asciiArt.ProcessArguments(args, asciiChars) // Process the first command-line argument with default ASCII characters
	}
}
