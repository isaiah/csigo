package main

import (
	"fmt"
	"testing"
)

func TestCalculatesCouplingByDegree(t *testing.T) {
	coupling := ByDegree(coupled) // coupled defined in coupling_algorithms_test.go
	assert(t, len(coupling), 3)
	assert(t, coupling[keyCombo{"A", "B"}], 1.0)
	assert(t, fmt.Sprintf("%.2f", coupling[keyCombo{"A", "C"}]), "0.67")
	assert(t, fmt.Sprintf("%.2f", coupling[keyCombo{"B", "C"}]), "0.67")
}

func TestCalculateCouplingByDegreeProvidingSingleEntity(t *testing.T) {
	singleEntity := []Entry{
		Entry{Prelude: &Prelude{Rev: "1"},
			Changes: []Change{Change{Entity: "this/is/a/single/entity"}}}}
	coupling := ByDegree(singleEntity)
	assert(t, coupling, make(map[keyCombo]float64))
}
