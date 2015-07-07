package main

// The idea behind effort is to identify how much each author
// contributed to a module. The measure here is a bit more
// rough than the churn metrics. On the other hand, the metric
// is available for all supported VCS.
// I use the generated statistics as a guide when refactoring;
// by ranking the authors based on their amount of contribution
// I know who to ask when visiting a new module.
//
// The analysis in the module is based on research by
// Marco D’Ambros, Harald C. Gall, Michele Lanza, and Martin Pinzger.

import (
	"math"
)

type authorRev struct {
	Author string
	Count  int
}

// Effort statistics how much work has put on a certain entity
type Effort struct {
	Entity        string
	Total         int
	Fragmentation float64
	AuthorRevs    []authorRev
}

// ByRevisionsPerAuthor returns all the contributors for the entity and their number of revisions
func ByRevisionsPerAuthor(changes []Change) (efforts []Effort) {
	entities, groupsByEntity := groupByEntity(changes)
	for _, entity := range entities {
		effort := Effort{Entity: entity}
		authors, groupsByAuthor := groupByAuthor(groupsByEntity[entity])
		for _, a := range authors {
			numRevs := len(groupsByAuthor[a])
			effort.AuthorRevs = append(effort.AuthorRevs, authorRev{Author: a, Count: numRevs})
			effort.Total = effort.Total + numRevs
		}
		efforts = append(efforts, effort)
	}
	return
}

// ByEntityFragmentation returns how distributed the effort on a specific entity is
func ByEntityFragmentation(changes []Change) (efforts []Effort) {
	efforts = ByRevisionsPerAuthor(changes)
	for i := range efforts {
		efforts[i].calculateFragmentation()
	}
	return
}

func (effort *Effort) calculateFragmentation() {
	var frag float64
	for _, rev := range effort.AuthorRevs {
		division := float64(rev.Count) / float64(effort.Total)
		frag += math.Pow(division, 2)
	}
	effort.Fragmentation = 1 - frag
}

// MainContributor returns the author and parcentage of revisions to the entity
func (effort *Effort) MainContributor() (string, float64) {
	var mainContrib *authorRev
	for _, rev := range effort.AuthorRevs {
		if mainContrib == nil || mainContrib.Count < rev.Count {
			mainContrib = &rev
		}
	}
	return mainContrib.Author, float64(mainContrib.Count) / float64(effort.Total)
}
