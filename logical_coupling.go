package main

// This module calculates the logical coupling of all modules.
//
// Logical coupling refers to modules that tend to change together.
// It's information that's recorded in our version-control systems (VCS).
//
// Input: all analysis expect an Incanter dataset with (at least) the following columns:
// :entity :rev
//
// Oputput: the analysis returns an Incanter dataset with the following columns:
// :entity :coupled :degree :average-revs

// ByDegree Calculates the degree of logical coupling. Returns a seq
// sorted in descending order (default) or an optional, custom sorting criterion.
// The calulcation is  based on the given coupling statistics.
// The coupling is calculated as a percentage value based on
// the number of shared commits between coupled entities divided
// by the average number of total commits for the coupled entities.
func ByDegree(entries []Entry) map[keyCombo]float64 {
	coupling := make(map[keyCombo]float64)
	numRevs := make(map[string]int)
	for _, change := range flatten(entries) {
		numRevs[change.Entity]++
	}
	for combo, count := range CouplingByRevision(entries) {
		average := float64(numRevs[combo.entity]+numRevs[combo.peer]) / 2
		coupling[combo] = float64(count) / average
	}
	return coupling
}
