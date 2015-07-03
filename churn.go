package main

import (
	"sort"
)

// Churn is the trend of the code
type Churn struct {
	// Dimension
	Date   string
	Author string
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
		for _, entry := range groups[date] {
			for _, change := range entry.Changes {
				churn.Added = churn.Added + change.LocAdded
				churn.Deleted = churn.Deleted + change.LocDeleted
			}
		}
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
		for _, entry := range groups[author] {
			for _, change := range entry.Changes {
				churn.Added = churn.Added + change.LocAdded
				churn.Deleted = churn.Deleted + change.LocDeleted
			}
		}
		churns = append(churns, churn)
	}
	return churns
}
