package main

import (
	"advent-of-code/internal"
	"fmt"
	"regexp"
	"strconv"
)

var part_regex = "[^a-zA-Z0-9 .]"
var part_regex_2 = "\\*"
var number_regex = "[\\d]+"

func main() {
	lines, err := internal.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var parts []Part
	for i, line := range lines {
		parts = append(parts, findParts(part_regex, i, line)...)
	}
	number_re := regexp.MustCompile(number_regex)
	var sum int
	for i, line := range lines {
		numbers := number_re.FindAllString(line, -1)
		locs := number_re.FindAllStringIndex(line, -1)
		for j, loc := range locs {
			if _, ok := checkNumberAdjacents(i, loc[0], loc[1]-1, parts); ok {
				num, err := strconv.Atoi(numbers[j])
				if err != nil {
					panic(err)
				}
				sum += num
			}
		}
	}

	fmt.Println(sum)
}

func part2(lines []string) {
	var parts []Part
	for i, line := range lines {
		parts = append(parts, findParts(part_regex_2, i, line)...)
	}
	number_re := regexp.MustCompile(number_regex)
	var sum int
	numbersByPart := make(map[Part][]int)
	for i, line := range lines {
		numbers := number_re.FindAllString(line, -1)
		locs := number_re.FindAllStringIndex(line, -1)
		for j, loc := range locs {
			if part, ok := checkNumberAdjacents(i, loc[0], loc[1]-1, parts); ok {
				num, err := strconv.Atoi(numbers[j])
				if err != nil {
					panic(err)
				}
				if len(numbersByPart[part]) == 0 {
					numbersByPart[part] = []int{num}
				} else {
					numbersByPart[part] = append(numbersByPart[part], num)
				}
			}
		}
	}
	for _, nums := range numbersByPart {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}

	fmt.Println(sum)
}

type Part struct {
	x, y int
}

func findParts(part_regex string, x int, input string) []Part {
	var parts []Part

	part_re := regexp.MustCompile(part_regex)
	locs := part_re.FindAllStringSubmatchIndex(input, -1)
	for _, loc := range locs {
		parts = append(parts, Part{x: x, y: loc[0]})
	}

	return parts
}

func checkNumberAdjacents(x, y1, y2 int, parts []Part) (Part, bool) {
	for _, part := range parts {
		if x-1 > part.x || x+1 < part.x {
			continue
		}
		if y1-1 > part.y || y2+1 < part.y {
			continue
		}
		return part, true
	}
	return Part{}, false
}
