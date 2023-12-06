package main

import (
	"advent-of-code/internal"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var card_regex = "Card *([0-9]+): "

func main() {
	lines, err := internal.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	card_re := regexp.MustCompile(card_regex)
	sum := 0
	for _, line := range lines {
		line = card_re.ReplaceAllString(line, "")
		winning, own := parseLine(line)
		sum += calculatePoints(winning, own)
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	card_re := regexp.MustCompile(card_regex)
	cards_count := make(map[int]int)
	for i := range lines {
		cards_count[i+1] += 1
	}
	for i, line := range lines {
		line = card_re.ReplaceAllString(line, "")
		winning, own := parseLine(line)
		including := checkContaining(winning, own)
		for j := i; j < i+including; j++ {
			if j+1 <= len(lines) {
				cards_count[j+2] += cards_count[i+1]
			}
		}
	}
	sum := 0
	for _, count := range cards_count {
		sum += count
	}
	fmt.Println(sum)
}

func parseLine(input string) ([]int, []int) {
	parts := strings.Split(input, " | ")
	var a []int
	var b []int
	convert := func(input string) []int {
		var out []int
		input = strings.ReplaceAll(strings.TrimPrefix(input, " "), "  ", " ")
		for _, val := range strings.Split(input, " ") {
			num, err := strconv.Atoi(strings.ReplaceAll(val, " ", ""))
			if err != nil {
				panic(err)
			}
			out = append(out, num)
		}
		return out
	}
	a = convert(parts[0])
	b = convert(parts[1])
	return a, b
}

func checkContaining(winning []int, own []int) int {
	found := 0
	for _, num := range winning {
		if internal.Contains[int](own, num) {
			found += 1
		}
	}
	return found
}

func calculatePoints(winning []int, own []int) int {
	found := checkContaining(winning, own)
	if found == 0 {
		return 0
	}
	return 1 * (int)(math.Pow(2, float64(found)-1.0))
}
