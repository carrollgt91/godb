package main

import (
	"fmt"
	"os"
)

// Entrypoint
func main() {
	ib := NewInputBuffer()
	fmt.Println("Running godb...")

	// Infinitely loop on user input
	for {
		// Read from buffer until we hit newline character
		text, err := ib.ReadLine()
		if err != nil {
			fmt.Printf("Error reading line: %v\n", err)
			os.Exit(ReadLineErrorCode)
		}

		// handle "meta" command
		if text[0] == '.' {
			err := DoMetaCommand(text)
			if err != nil {
				fmt.Printf("Error handling metacommand: %v\n", err)
			}

			continue
		}

		// handle SQL statements
		statement, err := PrepareStatement(text)

		if err != nil {
			fmt.Printf("Error preparing statement: %v\n", err)
			continue
		}

		fmt.Printf("Executing statement of type %v\n", statement.Type)
	}
}
