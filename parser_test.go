package main

import (
	"strings"
	"testing"
)

func TestParsePrelude(t *testing.T) {
	reader := strings.NewReader("--990442e--2013-08-29--Adam Petersen")
	parser := NewParser(reader)
	prelude, err := parser.Prelude()
	if err != nil {
		t.Fatal(err)
	}
	expected := Prelude{"Adam Petersen", "990442e", "2013-08-29"}
	if *prelude != expected {
		t.Errorf("expected %v, got %#v", expected, prelude)
	}
}

func TestParseChange(t *testing.T) {
	reader := strings.NewReader("1   0    project.clj")
	parser := NewParser(reader)
	change, err := parser.change()
	if err != nil {
		t.Fatal(err)
	}
	expected := Change{LocAdded: 1, LocDeleted: 0, Entry: "project.clj"}
	if expected != *change {
		t.Errorf("expected %v, got %v", expected, *change)
	}
}
