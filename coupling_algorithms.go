package main

type keyCombo struct {
	entity string
	peer   string
}

func CouplingByRevision(entries []Entry) map[keyCombo]int {
	coupling := make(map[keyCombo]int)
	for c := range entityCombinations(entries) {
		combo := keyCombo{c[0], c[1]}
		coupling[combo]++
	}
	return coupling
}

func entityCombinations(entries []Entry) <-chan []string {
	c := make(chan []string)
	go func() {
		for _, entry := range entries {
			changes := entry.Changes
			for i := 0; i < len(changes)-1; i++ {
				for j := i + 1; j < len(changes); j++ {
					c <- []string{changes[i].Entity, changes[j].Entity}
				}
			}
		}
		close(c)
	}()
	return c
}
