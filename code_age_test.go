package main

import (
	"fmt"
	"testing"
	"time"
)

var (
	day   = 24
	month = 30 * day

	commits = []Entry{
		Entry{Prelude: &Prelude{Rev: "1", Date: "2013-12-25"},
			Changes: []Change{
				Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "2", Date: "2013-12-31"},
			Changes: []Change{
				Change{Entity: "B"}}},
		Entry{Prelude: &Prelude{Rev: "3", Date: "2014-02-28"},
			Changes: []Change{
				Change{Entity: "A"}}},
		Entry{Prelude: &Prelude{Rev: "4", Date: "2014-04-05"},
			Changes: []Change{
				Change{Entity: "A"}}}}
)

func TestEntityAge(t *testing.T) {
	now, _ := time.Parse(DateLayout, "2014-04-06")
	threeMonths := 3 * month
	sixDays := 6 * day

	aAge, _ := time.ParseDuration("24h")
	bAge, _ := time.ParseDuration(fmt.Sprintf("%dh", threeMonths+sixDays))
	codeAge := ByAge(flatten(commits), now)
	assert(t, codeAge["A"], aAge)
	assert(t, codeAge["B"], bAge)
}

func TestCodeGetOlderAsTimePassesBy(t *testing.T) {
	// one month later
	now, _ := time.Parse(DateLayout, "2014-05-06")
	aAge, _ := time.ParseDuration(fmt.Sprintf("%dh", 1*month+day))
	bAge, _ := time.ParseDuration(fmt.Sprintf("%dh", 4*month+6*day))
	codeAge := ByAge(flatten(commits), now)
	assert(t, codeAge["A"], aAge)
	assert(t, codeAge["B"], bAge)

	// one year later
	now, _ = time.Parse(DateLayout, "2015-04-06")
	aAge, _ = time.ParseDuration(fmt.Sprintf("%dh", 12*month+6*day))
	bAge, _ = time.ParseDuration(fmt.Sprintf("%dh", 15*month+11*day))
	codeAge = ByAge(flatten(commits), now)
	assert(t, codeAge["A"], aAge)
	assert(t, codeAge["B"], bAge)

}
func TestCodeWasYoungerInThePast(t *testing.T) {
	now, _ := time.Parse(DateLayout, "2014-03-06")
	aAge, _ := time.ParseDuration(fmt.Sprintf("%dh", 6*day))
	bAge, _ := time.ParseDuration(fmt.Sprintf("%dh", 2*month+5*day))
	codeAge := ByAge(flatten(commits), now)
	assert(t, codeAge["A"], aAge)
	assert(t, codeAge["B"], bAge)
}
