package main

import (
	"fmt"
	"os"
	"strings"
)

/*
DoMetaCommand fully handles all the special, non-SQL commands.
For example, it will allow the user to exit the REPL by sending ".exit" to Stdin
*/
func DoMetaCommand(cmd string) error {
	if strings.Compare(".exit", cmd) == 0 {
		fmt.Println("Goodbye.")
		// exit with status code 0
		os.Exit(0)
		return nil
	}

	return fmt.Errorf("MetaCommand not found: %s", cmd)
}
