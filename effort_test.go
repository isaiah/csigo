package main

import (
	"reflect"
	"testing"
)

var (
	singleEffort = flatten(
		[]Entry{
			Entry{Prelude: &Prelude{Rev: "1", Author: "author", Date: "2015-07-01"},
				Changes: []Change{Change{Entity: "A"}}},
			Entry{Prelude: &Prelude{Rev: "2", Author: "author", Date: "2015-07-02"},
				Changes: []Change{Change{Entity: "B"}}},
			Entry{Prelude: &Prelude{Rev: "3", Author: "author", Date: "2015-07-03"},
				Changes: []Change{Change{Entity: "B"}}}})

	multipleEfforts = flatten(
		[]Entry{
			Entry{Prelude: &Prelude{Rev: "1", Author: "author1", Date: "2015-07-01"},
				Changes: []Change{Change{Entity: "A"}}},
			Entry{Prelude: &Prelude{Rev: "2", Author: "author2", Date: "2015-07-02"},
				Changes: []Change{Change{Entity: "A"}}},
			Entry{Prelude: &Prelude{Rev: "3", Author: "author3", Date: "2015-07-03"},
				Changes: []Change{Change{Entity: "A"}}}})
)

func TestEffortForSingleAuthor(t *testing.T) {
	efforts := ByRevisionsPerAuthor(singleEffort)
	if length := len(efforts); length != 2 {
		t.Errorf("expected 2 entities, got %d", length)
	}

	expected := Effort{
		Entity: "A", Total: 1,
		AuthorRevs: []authorRev{authorRev{"author", 1}}}

	if !reflect.DeepEqual(efforts[0], expected) {
		t.Errorf("expected first effort %v, got %v", expected, efforts[0])
	}

	expected = Effort{
		Entity: "B", Total: 2,
		AuthorRevs: []authorRev{authorRev{"author", 2}}}

	if !reflect.DeepEqual(efforts[1], expected) {
		t.Errorf("expected first effort %v, got %v", expected, efforts[1])
	}
}

func TestEffortsForMultipleAuthors(t *testing.T) {
	efforts := ByRevisionsPerAuthor(multipleEfforts)
	if len(efforts) != 1 {
		t.Fatalf("expected 1 entity, got %d", len(efforts))
	}
	expected := Effort{
		Entity: "A", Total: 3, AuthorRevs: []authorRev{
			authorRev{"author1", 1},
			authorRev{"author2", 1},
			authorRev{"author3", 1}}}
	effort := efforts[0]
	assert(t, effort, expected)
}

func assert(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("expected %v, got %v", a, b)
	}
}
