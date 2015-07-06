package main

import (
	"testing"
)

var (
	oneRev = []Entry{
		Entry{Prelude: &Prelude{Rev: "1"},
			Changes: []Change{
				Change{Entity: "A"},
				Change{Entity: "B"},
				Change{Entity: "C"}}}}
	coupled = []Entry{
		Entry{Prelude: &Prelude{Rev: "1"},
			Changes: []Change{
				Change{Entity: "A"},
				Change{Entity: "B"},
				Change{Entity: "C"}}},
		Entry{Prelude: &Prelude{Rev: "2"},
			Changes: []Change{
				Change{Entity: "A"},
				Change{Entity: "B"}}}}
)

func TestIdentifyCouplingInARevision(t *testing.T) {
	couplings := CouplingByRevision(oneRev)
	expected := map[keyCombo]int{keyCombo{"A", "B"}: 1, keyCombo{"A", "C"}: 1, keyCombo{"B", "C"}: 1}
	assert(t, couplings, expected)
}

func TestCouplingInMultipleRevisions(t *testing.T) {
	couplings := CouplingByRevision(coupled)
	expected := map[keyCombo]int{keyCombo{"A", "B"}: 2, keyCombo{"A", "C"}: 1, keyCombo{"B", "C"}: 1}
	assert(t, couplings, expected)
}
