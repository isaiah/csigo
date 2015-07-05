package main

import (
	"sort"
)

type EntityRevCount struct {
	Entity string
	Count  int
}

func ByRevision(changes []Change) (entities []EntityRevCount) {
	numOfRevs, entityGroups := sortByNumOfRev(groupEntitiesByRevCount(changes))
	for _, count := range numOfRevs {
		for _, entity := range entityGroups[count] {
			e := EntityRevCount{Entity: entity, Count: count}
			entities = append(entities, e)
		}
	}
	return
}

// groupEntitiesByRevCount returns the number of revision each entity has
func groupEntitiesByRevCount(changes []Change) map[string]int {
	groups := make(map[string]int)
	for _, change := range changes {
		groups[change.Entity] = groups[change.Entity] + 1
	}
	return groups
}

func sortByNumOfRev(entities map[string]int) (numOfRevs []int, entityGroups map[int][]string) {
	entityGroups = make(map[int][]string)
	for entity, count := range entities {
		entityGroups[count] = append(entityGroups[count], entity)
	}
	for count := range entityGroups {
		numOfRevs = append(numOfRevs, count)
	}
	// sort DESC
	sort.Sort(sort.Reverse(sort.IntSlice(numOfRevs)))
	return
}
