package main

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
		frag = frag + math.Pow(division, 2)
	}
	effort.Fragmentation = 1 - frag
}
