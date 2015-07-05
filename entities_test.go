package main

import (
	"testing"
)

var (
	vcsdTestData = []Entry{
		Entry{Prelude: &Prelude{Rev: "1", Author: "apt", Date: "2015-07-01"},
			Changes: []Change{
				Change{Entity: "A"},
				Change{Entity: "B"}}},
		Entry{Prelude: &Prelude{Rev: "2", Author: "apt", Date: "2015-07-01"},
			Changes: []Change{
				Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "1", Author: "jt", Date: "2015-07-01"},
			Changes: []Change{
				Change{Entity: "A"}}}}
)

func TestGroupEntityByRevCount(t *testing.T) {
	entities := ByRevision(flatten(vcsdTestData))
	if len(entities) != 2 {
		t.Fatalf("expected 2 entities, got %d", len(entities))
	}
	expected := EntityRevCount{"A", 3}
	if entities[0] != expected {
		t.Errorf("wrong entity count, expected %v, got %v", expected, entities[0])
	}
	expected = EntityRevCount{"B", 1}
	if entities[1] != expected {
		t.Errorf("wrong entity count, expected %v, got %v", expected, entities[1])
	}
}
