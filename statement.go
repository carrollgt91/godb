package main

import (
	"fmt"
	"strings"
)

/*
StatementType is an enum-like to represent what type of SQL statement
the given Statement is representing.
*/
type StatementType int

const (
	StatementInsert = 0
	StatementSelect = 1
)

/*
Statement represents a prepared SQL statement that will be compiled into
bytecode after parameters have been extracted.
*/
type Statement struct {
	Type StatementType
}

/*
PrepareStatement takes the raw text passed by a user query and parses it
into a statement, extracting relevant parameters and storing them in the
returned Statement.
*/
func PrepareStatement(text string) (*Statement, error) {
	lcaseText := strings.ToLower(text)
	if strings.Contains(lcaseText, "insert") {
		return &Statement{
			StatementInsert,
		}, nil
	}

	if strings.Contains(lcaseText, "select") {
		return &Statement{
			StatementSelect,
		}, nil
	}

	return nil, fmt.Errorf("Unrecognized statement %v", text)
}
