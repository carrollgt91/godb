package main

import (
	"fmt"
	"os"
	"strings"
)

// Entrypoint
func main() {
	ib := NewInputBuffer()
	fmt.Println("Running godb...")

	for {
		// Read from buffer until we hit newline character
		text, err := ib.ReadLine()
		if err != nil {
			fmt.Printf("Error reading line: %v\n", err)
			os.Exit(ReadLineErrorCode)
		}

		if strings.Compare(".exit", text) == 0 {
			fmt.Println("Goodbye.")
			// exit with status code 0
			os.Exit(0)
			return
		}

		fmt.Printf("Error: Command Not recognized %s\n", text)
	}
}
