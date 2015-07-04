package main

import (
	"testing"
)

var (
	simple = flatten([]Entry{
		Entry{Prelude: &Prelude{Rev: "1", Date: "2015-07-01", Author: "IS"},
			Changes: []Change{
				Change{LocAdded: 10, LocDeleted: 1, Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "2", Date: "2015-07-02", Author: "SI"},
			Changes: []Change{
				Change{LocAdded: 20, LocDeleted: 2, Entity: "B"}}},
		Entry{Prelude: &Prelude{Rev: "3", Date: "2015-07-02", Author: "IS"},
			Changes: []Change{
				Change{LocAdded: 2, LocDeleted: 0, Entity: "B"}}}})
)

func TestAbsolutesTrend(t *testing.T) {
	churns := AbsoluteTrends(simple)
	if len(churns) != 2 {
		t.Fatalf("expected 2 churns, got %d", len(churns))
	}
	if churns[0].Date != "2015-07-01" {
		t.Error("churns are not sorted")
	}
	if churns[0].Added != 10 || churns[0].Deleted != 1 {
		t.Errorf("wrong Added %d, Deleted %d", churns[0].Added, churns[0].Deleted)
	}
}

func TestChurnByAuthor(t *testing.T) {
	churns := ByAuthor(simple)
	if len(churns) != 2 {
		t.Fatalf("expected 2 churns, got %d", len(churns))
	}
	if churns[0].Author != "IS" {
		t.Error("churns are not sorted")
	}
	if churns[0].Added != 12 || churns[0].Deleted != 1 {
		t.Errorf("wrong Added %d, Deleted %d", churns[0].Added, churns[0].Deleted)
	}
}

func TestChurnByEntity(t *testing.T) {
	churns := ByEntity(simple)
	if len(churns) != 2 {
		t.Fatalf("expected 2 churns, got %d", len(churns))
	}
	if churns[0].Entity != "A" {
		t.Error("churns are not sorted")
	}
	if churns[0].Added != 10 || churns[0].Deleted != 1 {
		t.Errorf("wrong Added %d, Deleted %d", churns[0].Added, churns[0].Deleted)
	}
}

func TestChurnByOwnership(t *testing.T) {
	churns := ByOwnership(simple)
	if len(churns) != 3 {
		t.Fatalf("expected 3 churns, got %d", len(churns))
	}
}
