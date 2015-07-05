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

	sameAuthor = flatten([]Entry{
		Entry{Prelude: &Prelude{Rev: "1", Date: "2015-07-01", Author: "author1"},
			Changes: []Change{
				Change{LocAdded: 10, LocDeleted: 1, Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "2", Date: "2015-07-02", Author: "author1"},
			Changes: []Change{
				Change{LocAdded: 2, LocDeleted: 5, Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "3", Date: "2015-07-02", Author: "author2"},
			Changes: []Change{
				Change{LocAdded: 7, LocDeleted: 1, Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "4", Date: "2015-07-02", Author: "author2"},
			Changes: []Change{
				Change{LocAdded: 8, LocDeleted: 2, Entity: "A"}}}})

	singleAuthor = flatten([]Entry{
		Entry{Prelude: &Prelude{Rev: "1", Date: "2015-07-01", Author: "single"},
			Changes: []Change{
				Change{LocAdded: 10, LocDeleted: 1, Entity: "Same"}}},
		Entry{Prelude: &Prelude{Rev: "1", Date: "2015-07-02", Author: "single"},
			Changes: []Change{
				Change{LocAdded: 20, LocDeleted: 2, Entity: "Same"}}},
		Entry{Prelude: &Prelude{Rev: "1", Date: "2015-07-02", Author: "single"},
			Changes: []Change{
				Change{LocAdded: 2, LocDeleted: 0, Entity: "Same"}}}})
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
	expectedChurn := Churn{Entity: "A", Author: "IS", Added: 10, Deleted: 1}
	if churns[0] != expectedChurn {
		t.Errorf("wrong churn, expected %v, got %v", expectedChurn, churns[0])
	}
	expectedChurn = Churn{Entity: "B", Author: "IS", Added: 2, Deleted: 0}
	if churns[1] != expectedChurn {
		t.Errorf("wrong churn, expected %v, got %v", expectedChurn, churns[1])
	}
	expectedChurn = Churn{Entity: "B", Author: "SI", Added: 20, Deleted: 2}
	if churns[2] != expectedChurn {
		t.Errorf("wrong churn, expected %v, got %v", expectedChurn, churns[2])
	}
}

func TestChurnByOwnershipForSameAuthor(t *testing.T) {
	churns := ByOwnership(sameAuthor)
	if len(churns) != 2 {
		t.Fatalf("expected 2 churns, got %d", len(churns))
	}
	expectedChurn := Churn{Entity: "A", Author: "author1", Added: 12, Deleted: 6}
	if churns[0] != expectedChurn {
		t.Errorf("wrong churn, expected %v, got %v", expectedChurn, churns[0])
	}
	expectedChurn = Churn{Entity: "A", Author: "author2", Added: 15, Deleted: 3}
	if churns[1] != expectedChurn {
		t.Errorf("wrong churn, expected %v, got %v", expectedChurn, churns[1])
	}
}

func TestIdentifySingleMainContributor(t *testing.T) {
	mainDev := ByMainContributor(singleAuthor)
	expected := Contributor{Entity: "Same", Author: "single", Added: 32, Total: 32}
	if len(mainDev) != 1 {
		t.Fatalf("expected only one contributor, got %d", len(mainDev))
	}
	if mainDev[0] != expected {
		t.Errorf("expected main contributor %v, got %v", expected, mainDev[0])
	}
	if mainDev[0].Ownership() != 1 {
		t.Errorf("Ownership of a single developer should be 1.0, got %f", mainDev[0].Ownership())
	}
}

func TestIdentifyMainContributorOnSharedEntities(t *testing.T) {
	mainDevs := ByMainContributor(sameAuthor)
	expected := Contributor{Entity: "A", Author: "author2", Added: 15, Total: 27}
	if mainDevs[0] != expected {
		t.Errorf("expected main contributor %v, got %v", expected, mainDevs[0])
	}
}
