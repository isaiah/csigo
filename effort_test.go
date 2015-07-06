package main

import (
	"fmt"
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

func TestEntityFragmentationForSingleEffort(t *testing.T) {
	efforts := ByEntityFragmentation(singleEffort)
	assert(t, efforts[0].Entity, "A")
	assert(t, efforts[0].Fragmentation, 0.0)
	assert(t, efforts[1].Entity, "B")
	assert(t, efforts[1].Fragmentation, 0.0)
}

func TestEntityFragmentationForMultipleEfforts(t *testing.T) {
	efforts := ByEntityFragmentation(multipleEfforts)
	assert(t, efforts[0].Entity, "A")
	assert(t, fmt.Sprintf("%.2f", efforts[0].Fragmentation), "0.67")
}

func TestIdentifyMainContributorByRevisions(t *testing.T) {
	sharedEffort := []Entry{
		Entry{Prelude: &Prelude{Rev: "1", Author: "a1", Date: "2015-07-01"},
			Changes: []Change{Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "2", Author: "a2", Date: "2015-07-02"},
			Changes: []Change{Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "3", Author: "a2", Date: "2015-07-03"},
			Changes: []Change{Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "4", Author: "a3", Date: "2015-07-04"},
			Changes: []Change{Change{Entity: "B"}}},
		Entry{Prelude: &Prelude{Rev: "5", Author: "a4", Date: "2015-07-05"},
			Changes: []Change{Change{Entity: "C"}}},
		Entry{Prelude: &Prelude{Rev: "6", Author: "a5", Date: "2015-07-06"},
			Changes: []Change{Change{Entity: "C"}}},
	}

	efforts := ByRevisionsPerAuthor(flatten(sharedEffort))

	if len(efforts) != 3 {
		t.Fatalf("expected 3 efforts, got %d", len(efforts))
	}
	author, ownership := efforts[0].MainContributor()
	if author != "a2" || fmt.Sprintf("%.2f", ownership) != "0.67" {
		t.Errorf("main contributor to the first entity should be %s, got %s", "a2", author)
	}
	author, ownership = efforts[1].MainContributor()
	if author != "a3" || ownership != 1 {
		t.Errorf("main contributor to the second entity should be %s, got %s", "a3", author)
	}
	author, ownership = efforts[2].MainContributor()
	if author != "a5" || fmt.Sprintf("%.2f", ownership) != "0.50" {
		t.Errorf("main contributor to the third entity should be %s, got %s", "a5", author)
	}

}

func assert(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("expected %v, got %v", b, a)
	}
}
