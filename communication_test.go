package main

import (
	"testing"
)

func TestCombination(t *testing.T) {
	data := []authorRev{
		authorRev{Author: "A"},
		authorRev{Author: "B"},
		authorRev{Author: "C"}}
	c := combinations([]Effort{Effort{AuthorRevs: data}})
	var combos [][]string
	for combo := range c {
		combos = append(combos, combo)
	}
	assert(t, combos, [][]string{[]string{"A", "B"}, []string{"A", "C"}, []string{"B", "C"}})
}

func TestCommunicationNeedsForSharedAuthorship(t *testing.T) {
	sharingAuthors := []Entry{
		Entry{Prelude: &Prelude{Rev: "1", Author: "at", Date: "2015-07-01"},
			Changes: []Change{Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "2", Author: "jt", Date: "2015-07-02"},
			Changes: []Change{Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "3", Author: "ap", Date: "2015-07-03"},
			Changes: []Change{Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "4", Author: "at", Date: "2015-07-04"},
			Changes: []Change{Change{Entity: "B"}}},
		Entry{Prelude: &Prelude{Rev: "5", Author: "jt", Date: "2015-07-05"},
			Changes: []Change{Change{Entity: "B"}}}}
	coms := BySharedEntities(Flatten(sharingAuthors))
	assert(t, len(coms), 3)
	for _, com := range coms {
		if com.Author == "at" && com.Peer == "jt" {
			assert(t, com, Communication{Author: "at", Peer: "jt", Shared: 2, Average: 2, Strength: 1})
		} else if com.Author == "ap" && com.Peer == "jt" {
			assert(t, com, Communication{Author: "ap", Peer: "jt", Shared: 1, Average: 2, Strength: 0.5})
		} else {
			assert(t, com, Communication{Author: "ap", Peer: "at", Shared: 1, Average: 2, Strength: 0.5})
		}
	}
}
