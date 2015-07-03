package main

import (
	"strings"
	"testing"
)

var (
	expectedPrelude = Prelude{"Adam Petersen", "990442e", "2013-08-29"}
	expectedChange  = Change{LocAdded: 1, LocDeleted: 0, Entry: "project.clj"}
	expectedChange2 = Change{LocAdded: 2, LocDeleted: 4, Entry: "src/code_maat/parsers/git.clj"}
	expectedEntry   = Entry{Prelude: &expectedPrelude, Changes: []Change{expectedChange, expectedChange2}}
)

func TestParsePrelude(t *testing.T) {
	reader := strings.NewReader("--990442e--2013-08-29--Adam Petersen")
	parser := NewParser(reader)
	prelude, err := parser.prelude()
	if err != nil {
		t.Fatal(err)
	}
	if *prelude != expectedPrelude {
		t.Errorf("expected %v, got %#v", expectedPrelude, prelude)
	}
}

func TestParseChange(t *testing.T) {
	reader := strings.NewReader("1   0    project.clj")
	parser := NewParser(reader)
	change, err := parser.change()
	if err != nil {
		t.Fatal(err)
	}
	if expectedChange != *change {
		t.Errorf("expected %v, got %v", expectedChange, *change)
	}
}

func TestParseEntry(t *testing.T) {
	commit := `
--990442e--2013-08-29--Adam Petersen
1    0    project.clj
2    4    src/code_maat/parsers/git.clj
`
	reader := strings.NewReader(commit)
	parser := NewParser(reader)
	entry, err := parser.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if *entry.Prelude != *expectedEntry.Prelude {
		t.Errorf("unmatch prelude %v, got %v", *expectedEntry.Prelude, *entry.Prelude)
	}
	if len(entry.Changes) != len(expectedEntry.Changes) {
		t.Fatalf("unmatch changes %v, got %v", expectedEntry.Changes, entry.Changes)
	}
	for i, change := range expectedEntry.Changes {
		if change != entry.Changes[i] {
			t.Errorf("change doens't match %v, got %v", change, entry.Changes[i])
		}
	}
}
