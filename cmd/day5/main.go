package main

import (
	"advent-of-code/internal"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := internal.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	seeds := strings.Split(strings.ReplaceAll(lines[0], "seeds: ", ""), " ")
	var seedRanges []seedRange
	for _, seed := range seeds {
		val, err := strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seedRanges = append(seedRanges, seedRange{start: val, stop: val})
	}
	mappedValues := parseInput(seedRanges, lines)
	sort.Ints(mappedValues)
	fmt.Println(mappedValues[0])
}

func part2(lines []string) {
	seeds := strings.Split(strings.ReplaceAll(lines[0], "seeds: ", ""), " ")
	var seedRanges []seedRange
	for i := 0; i < len(seeds); i += 2 {
		start, err := strconv.Atoi(seeds[i])
		if err != nil {
			panic(err)
		}
		seedsRange, err := strconv.Atoi(seeds[i+1])
		if err != nil {
			panic(err)
		}
		seedRanges = append(seedRanges, seedRange{start: start, stop: start + seedsRange - 1})
	}
	mappedValues := parseInput(seedRanges, lines)
	sort.Ints(mappedValues)
	fmt.Println(mappedValues[0])
}

func parseInput(seedRanges []seedRange, lines []string) []int {
	var conversions []conversion
	for i := 2; i < len(lines); i++ {
		line := lines[i]
		if line == "" {
			var temp []seedRange
			for _, seedRange := range seedRanges {
				intersection := seedRange.intersect(conversions)
				temp = append(temp, intersection...)
			}
			seedRanges = temp
			conversions = []conversion{}
			continue
		}
		if strings.Contains(line, "map") {
			continue
		}
		elems := strings.Split(line, " ")
		destStart, err := strconv.Atoi(elems[0])
		if err != nil {
			panic(err)
		}
		sourceStart, err := strconv.Atoi(elems[1])
		if err != nil {
			panic(err)
		}
		itemRange, err := strconv.Atoi(elems[2])
		if err != nil {
			panic(err)
		}
		conversions = append(conversions, conversion{sourceStart: sourceStart, destStart: destStart, itemRange: itemRange})
	}
	var temp []seedRange
	for _, seedRange := range seedRanges {
		intersection := seedRange.intersect(conversions)
		temp = append(temp, intersection...)
	}
	result := make([]int, len(temp))
	for i, seedRange := range temp {
		result[i] = seedRange.start
	}
	return result
}

type seedRange struct {
	start, stop int
}

func (s seedRange) intersect(conversions []conversion) []seedRange {
	var out []seedRange
	for _, conversion := range conversions {
		sourceStart := conversion.sourceStart
		destStart := conversion.destStart
		sourceStop := sourceStart + conversion.itemRange - 1
		destStop := destStart + conversion.itemRange - 1

		if s.start > sourceStop || s.stop < sourceStart {
			continue
		} else if s.start >= sourceStart {
			if s.stop <= sourceStop {
				tempStart := destStart + s.start - sourceStart
				out = append(out, seedRange{start: destStart + s.start - sourceStart, stop: tempStart + s.stop - s.start})
				return out
			} else if s.stop > sourceStop {
				out = append(out, seedRange{start: destStart + s.start - sourceStart, stop: destStop})
				s = seedRange{start: sourceStop + 1, stop: s.stop}
			}
		} else {
			if s.stop <= sourceStop {
				out = append(out, seedRange{start: destStart, stop: destStart + s.stop - sourceStart})
				s = seedRange{start: s.start, stop: sourceStart - 1}
			} else if s.stop > sourceStop {
				out = append(out, seedRange{start: s.start, stop: sourceStart - 1}.intersect(conversions)...)
				out = append(out, seedRange{start: destStart, stop: destStop})
				out = append(out, seedRange{start: sourceStop + 1, stop: s.stop}.intersect(conversions)...)
			}
		}
	}
	out = append(out, s)
	return out
}

type conversion struct {
	sourceStart, destStart, itemRange int
}
