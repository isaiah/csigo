package main

type authorRev struct {
	Author string
	Count  int
}
type Effort struct {
	Entity     string
	Total      int
	AuthorRevs []authorRev
}

func ByRevisionsPerAuthor(changes []Change) (efforts []Effort) {
	entities, groupsByEntity := groupByEntity(changes)
	for _, entity := range entities {
		effort := Effort{Entity: entity}
		authors, groupsByAuthor := groupByAuthor(groupsByEntity[entity])
		for _, a := range authors {
			numRevs := len(groupsByAuthor[a])
			effort.AuthorRevs = append(effort.AuthorRevs, authorRev{Author: a, Count: numRevs})
			effort.Total = effort.Total + numRevs
		}
		efforts = append(efforts, effort)
	}
	return
}
