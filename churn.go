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

func AbsoluteTrends(changes []Change) (churns []Churn) {
	dates, groups := groupByDate(changes)
	for _, date := range dates {
		churn := Churn{Date: date}
		churn.sumByChanges(groups[date])
		churns = append(churns, churn)
	}
	return churns
}

// ByAuthor calculates the churn trend by author
func ByAuthor(changes []Change) (churns []Churn) {
	authors, groups := groupByAuthor(changes)
	for _, author := range authors {
		churn := Churn{Author: author}
		churn.sumByChanges(groups[author])
		churns = append(churns, churn)
	}
	return
}

func (c *Churn) sumByChanges(changes []Change) {
	for _, change := range changes {
		c.Added = c.Added + change.LocAdded
		c.Deleted = c.Deleted + change.LocDeleted
	}
}

// ByEntity calculates churns by files
func ByEntity(changes []Change) (churns []Churn) {
	entities, groups := groupByEntity(changes)
	for _, entity := range entities {
		churn := Churn{Entity: entity}
		churn.sumByChanges(groups[entity])
		churns = append(churns, churn)
	}
	return
}

// ByOwnership returns a table specifying the ownership of each module.
// Ownership is defined as the amount of churn contributed by each author to each entity.
func ByOwnership(changes []Change) (churns []Churn) {
	// group by entity first
	entities, groupsByEntity := groupByEntity(changes)
	for _, entity := range entities {
		authors, groupsByAuthor := groupByAuthor(groupsByEntity[entity])
		for _, author := range authors {
			churn := Churn{Author: author, Entity: entity}
			churn.sumByChanges(groupsByAuthor[author])
			churns = append(churns, churn)
		}
	}
	return
}

func groupByEntity(changes []Change) (entities []string, groups map[string][]Change) {
	groups = make(map[string][]Change)
	for _, change := range changes {
		groups[change.Entity] = append(groups[change.Entity], change)
	}
	entities = stringKeys(groups)
	return
}

func groupByAuthor(changes []Change) (authors []string, groups map[string][]Change) {
	groups = make(map[string][]Change)
	for _, change := range changes {
		groups[change.Author] = append(groups[change.Author], change)
	}
	authors = stringKeys(groups)
	return
}

func groupByDate(changes []Change) (dates []string, groups map[string][]Change) {
	groups = make(map[string][]Change)
	for _, change := range changes {
		groups[change.Date] = append(groups[change.Date], change)
	}
	dates = stringKeys(groups)
	return
}

func stringKeys(groups map[string][]Change) (keys []string) {
	for key := range groups {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return
}
