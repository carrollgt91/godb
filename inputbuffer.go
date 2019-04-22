package main

import (
	"bufio"
	"os"
	"strings"
)

/*
InputBuffer is a wrapper around taking input from the user via
the command prompt.
*/
type InputBuffer struct {
	reader *bufio.Reader
}

/*
NewInputBuffer creates a new InputBuffer that reads from Stdin
*/
func NewInputBuffer() *InputBuffer {
	ib := &InputBuffer{
		reader: bufio.NewReader(os.Stdin),
	}

	return ib
}

/*
ReadLine reads a single line of user input and returns it, or an error if one exists.
*/
func (ib *InputBuffer) ReadLine() (string, error) {
	text, err := ib.reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	// Remove trailing newline
	text = strings.Replace(text, "\n", "", -1)
	return text, nil
}
