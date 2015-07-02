package main

import (
	"strings"
	"testing"
)

const (
	entry = `--990442e--2013-08-29--Adam Petersen
1⇥0⇥project.clj
2⇥4⇥src/code_maat/parsers/git.clj
`
)

func TestParseSingleEntry(t *testing.T) {
	parser := NewParser(strings.NewReader(entry))
	prelude, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	expected := Prelude{"990442e", "Adam Petersen", "2013-08-29"}
	if *prelude != expected {
		t.Errorf("expected %v, got %#v", expected, prelude)
	}
}
