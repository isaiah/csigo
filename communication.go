package main

import (
	"math"
)

//  This module attempts to give some heuristics on
//  the communication needs of a project.
//  The idea is basedo on Conway's law - a project
//  works best when its organizational structure is
//  mirrored in software.
//
//  The algorithm is similiar to the one used for
//  logical coupling: calculate the number of shared
//  commits between all permutations of authors.
//  Based on their total averaged commits, a
//  communication strength value is calculated.

// Communication frequencies of two contributors co-authoring
type Communication struct {
	Author   string
	Peer     string
	Shared   int
	Average  float64
	Strength float64
}

// BySharedEntities calculates communications between two contributors, based on how many entities they co-authored
func BySharedEntities(changes []Change) (communications []Communication) {
	// XXX fix the calculation, divide occurance of co-author (num ofenity) by average of total commit is wrong
	efforts := ByRevisionsPerAuthor(changes)
	perAuthor := revsPerAuthor(efforts)
	comboCount := make(map[keyCombo]int)
	for combo := range combinations(efforts) {
		names := keyCombo{combo[0], combo[1]}
		comboCount[names]++
	}
	for c, count := range comboCount {
		initial, peer := c.entity, c.peer
		average := math.Ceil(float64(perAuthor[initial]+perAuthor[peer]) / 2)
		communication := Communication{Author: initial, Peer: peer, Shared: count,
			Average: average, Strength: float64(count) / average}
		communications = append(communications, communication)
	}
	return
}

// get combinations of authors co-contribute the entity
func combinations(efforts []Effort) <-chan []string {
	c := make(chan []string)
	go func() {
		for _, effort := range efforts {
			revs := effort.AuthorRevs
			for i := 0; i < len(revs)-1; i++ {
				for j := i + 1; j < len(revs); j++ {
					c <- []string{revs[i].Author, revs[j].Author}
				}
			}
		}
		close(c)
	}()
	return c
}

func revsPerAuthor(efforts []Effort) map[string]int {
	perAuthor := make(map[string]int)
	for _, effort := range efforts {
		for _, author := range effort.AuthorRevs {
			perAuthor[author.Author] += author.Count
		}
	}
	return perAuthor
}
