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
	commit := `--990442e--2013-08-29--Adam Petersen
1    0    project.clj
2    4    src/code_maat/parsers/git.clj
`
	reader := strings.NewReader(commit)
	parser := NewParser(reader)
	entry, err := parser.entry()
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

func TestParseEntries(t *testing.T) {
	logs := `--b777738--2013-08-29--Adam Petersen
10	9	src/code_maat/parsers/git.clj
32	0	test/code_maat/parsers/git_test.clj

--a527b79--2013-08-29--Adam Petersen
6	2	src/code_maat/parsers/git.clj
0	7	test/code_maat/end_to_end/scenario_tests.clj
18	0	test/code_maat/end_to_end/simple_git.txt
21	0	test/code_maat/end_to_end/svn_live_data_test.clj
`
	reader := strings.NewReader(logs)
	parser := NewParser(reader)
	entries, _ := parser.Parse()
	if len(entries) != 2 {
		t.Fatalf("wrong number of entries found %d", len(entries))
	}
	if entries[1].Changes[3].LocAdded != 21 {
		t.Error("the forth change of the second commit has 21 LocAdded")
	}
}
