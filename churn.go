package main

import (
	"sort"
)

// Churn is the trend of the code
type Churn struct {
	// Dimension
	Date   string
	Author string
	Entity string
	// Measure
	Added   int64
	Deleted int64
}

func AbsoluteTrends(entries []Entry) []Churn {
	var churns []Churn
	groups := groupByDate(entries)
	var keys []string
	for date := range groups {
		keys = append(keys, date)
	}
	sort.Strings(keys)
	for _, date := range keys {
		churn := Churn{Date: date}
		churn.totalChurn(groups[date])
		churns = append(churns, churn)
	}
	return churns
}

func groupByDate(entries []Entry) map[string][]Entry {
	ret := make(map[string][]Entry)
	for _, entry := range entries {
		ret[entry.Date] = append(ret[entry.Date], entry)
	}
	return ret
}

// ByAuthor calculates the churn trend by author
func ByAuthor(entries []Entry) []Churn {
	churns := []Churn{}
	groups := make(map[string][]Entry)
	for _, entry := range entries {
		groups[entry.Author] = append(groups[entry.Author], entry)
	}
	var authors []string
	for author := range groups {
		authors = append(authors, author)
	}
	sort.Strings(authors)
	for _, author := range authors {
		churn := Churn{Author: author}
		churn.totalChurn(groups[author])
		churns = append(churns, churn)
	}
	return churns
}

func (c *Churn) totalChurn(entries []Entry) {
	for _, entry := range entries {
		c.updateByChanges(entry.Changes)
	}
}

func (c *Churn) updateByChanges(changes []Change) {
	for _, change := range changes {
		c.Added = c.Added + change.LocAdded
		c.Deleted = c.Deleted + change.LocDeleted
	}
}

// ByEntity calculates churns by files
func ByEntity(entries []Entry) []Churn {
	churns := []Churn{}
	groups := make(map[string][]Change)
	for _, entry := range entries {
		for _, change := range entry.Changes {
			groups[change.Entity] = append(groups[change.Entity], change)
		}
	}
	var entities []string
	for entity := range groups {
		entities = append(entities, entity)
	}
	sort.Strings(entities)
	for _, entity := range entities {
		churn := Churn{Entity: entity}
		churn.updateByChanges(groups[entity])
		churns = append(churns, churn)
	}
	return churns
}
