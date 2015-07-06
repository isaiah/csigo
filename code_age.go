package main

import (
	"time"
)

var (
	// DateLayout the layout of the date format given by git log
	DateLayout = "2006-01-02"
)

// ByAge returns the age of each entity compared to the provided time
func ByAge(changes []Change, until time.Time) map[string]time.Duration {
	var latestChanges []*Change
	now := until.Format(DateLayout)
	entities, groups := groupByEntity(changes)
	for _, entity := range entities {
		cs := groups[entity]
		var latestChange *Change
		for i := range groups[entity] {
			change := cs[i]
			if (latestChange == nil || latestChange.Date <= change.Date) && change.Date < now {
				latestChange = &change
			}
		}
		if latestChange != nil {
			latestChanges = append(latestChanges, latestChange)
		}
	}
	entityAge := make(map[string]time.Duration)
	for _, change := range latestChanges {
		if mtime, err := time.Parse(DateLayout, change.Date); err == nil {
			entityAge[change.Entity] = until.Sub(mtime)
		}
	}
	return entityAge
}
