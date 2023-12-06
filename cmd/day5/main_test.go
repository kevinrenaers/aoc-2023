package main

import (
	"fmt"
	"testing"
)

func Test_parseInput(t *testing.T) {
	a := seedRange{start: 5, stop: 10}
	c1 := conversion{sourceStart: 1, destStart: 4, itemRange: 6}
	c2 := conversion{sourceStart: 1, destStart: 4, itemRange: 3}
	c3 := conversion{sourceStart: 10, destStart: 9, itemRange: 1}
	ranges := a.intersect([]conversion{c1, c2, c3})
	fmt.Println(ranges)
	// []seedRange{{start: 8, stop: 10}, {start: 8, stop: 10}
	a = seedRange{start: 5, stop: 10}
	c1 = conversion{sourceStart: 10, destStart: 4, itemRange: 1}
	ranges = a.intersect([]conversion{c1})
	fmt.Println(ranges)
	// assert.Equal(t, ranges, []seedRange{{start: 5, stop: 9}, {start: 4, stop: 4}})
	a = seedRange{start: 5, stop: 10}
	c1 = conversion{sourceStart: 1, destStart: 4, itemRange: 4}
	ranges = a.intersect([]conversion{c1})
	fmt.Println(ranges)
	// assert.Equal(t, ranges, []seedRange{{start: 8, stop: 8}, {start: 6, stop: 10}})
	a = seedRange{start: 5, stop: 10}
	c1 = conversion{sourceStart: 6, destStart: 4, itemRange: 2}
	ranges = a.intersect([]conversion{c1})
	fmt.Println(ranges)
	// assert.Equal(t, ranges, []seedRange{{start: 5, stop: 5}, {start: 4, stop: 6}, {start: 8, stop: 10}})
}
