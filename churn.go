package main

import (
	"sort"
)

// Churn is the trend of the code
type Churn struct {
	Date    string
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
		entries := groups[date]
		for _, entry := range entries {
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
