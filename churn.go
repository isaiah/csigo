package main

// This module contains functions for calculating churn metrics.
// Code churn is related to the quality of modules; the higher
// the churn, the more post-release defects.
// Further, inspecting the churn trend lets us spot certain
// organization-oriented patterns. For example, we may spot
// integration bottlenecks as spikes just before the end of
// one iteration.

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

// Contributor is a author of the entity
type Contributor struct {
	Entity string
	Author string
	Added  int64
	Total  int64
}

func (c *Churn) sumByChanges(changes []Change) {
	for _, change := range changes {
		c.Added += change.LocAdded
		c.Deleted += change.LocDeleted
	}
}

func (contrib *Contributor) Ownership() float64 {
	return float64(contrib.Added) / float64(contrib.Total)
}

// ByMainContributor Identify the main contributor of each entity, main contributor is the author of the entity who contributed the most lines of codes
func ByMainContributor(changes []Change) (cs []Contributor) {
	churnGroups := groupByAuthorContribution(changes)
	for entity, churns := range churnGroups {
		var totalContrib int64
		var mainContrib *Churn
		for _, churn := range churns {
			if mainContrib == nil || mainContrib.Added < churn.Added {
				mainContrib = &churn
			}
			totalContrib += churn.Added
		}
		contributor := Contributor{Entity: entity, Author: mainContrib.Author, Added: mainContrib.Added, Total: totalContrib}
		cs = append(cs, contributor)
	}
	return
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
	groupsByAuthorContrib := groupByAuthorContribution(changes)
	for _, cs := range groupsByAuthorContrib {
		churns = append(churns, cs...)
	}
	return
}

func groupByAuthorContribution(changes []Change) map[string][]Churn {
	groupsByAuthorContrib := make(map[string][]Churn)
	// group by entity first
	entities, groupsByEntity := groupByEntity(changes)
	for _, entity := range entities {
		authors, groupsByAuthor := groupByAuthor(groupsByEntity[entity])
		for _, author := range authors {
			churn := Churn{Author: author, Entity: entity}
			churn.sumByChanges(groupsByAuthor[author])
			groupsByAuthorContrib[entity] = append(groupsByAuthorContrib[entity], churn)
		}
	}
	return groupsByAuthorContrib
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
