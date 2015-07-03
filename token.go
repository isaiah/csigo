package main

import (
	"regexp"
)

// Token represents a lexical token
type Token int

var (
	rev      = regexp.MustCompile(`^[\da-f]+$`)
	author   = regexp.MustCompile(`^[\w\s]+?$`)
	date     = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	numstate = regexp.MustCompile(`\d+?$`)
	file     = regexp.MustCompile(`^[\/\w]*\.\w+$`)
)

const (
	// Special tokens
	ILLEGAL Token = iota
	EOF           // End of file
	WS            // white space
	TAB
	NL

	// Literal
	ENTRY // main
	PRELUDE
	REV
	AUTHOR
	DATE
	CHANGES
	CHANGE
	ADDED
	DELETED
	NUMSTAT
	FILE
	SEPARATOR

	// Keywords
)

// Prelude represents the prelude of a git entry
type Prelude struct {
	Author string
	Rev    string
	Date   string
}

// Change is the entry of change
type Change struct {
	LocAdded   int64
	LocDeleted int64
	Entity     string // file name
}

// Entry is a commit
type Entry struct {
	*Prelude
	Changes []Change
}
