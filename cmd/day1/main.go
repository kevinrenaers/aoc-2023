package main

import (
	"advent-of-code/internal"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

var regex = "[0-9]|one|two|three|four|five|six|seven|eight|nine"

func main() {
	lines, err := internal.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	re := regexp.MustCompile("[0-9]{1}")
	sum := 0
	for _, line := range lines {
		digits := re.FindAllString(line, -1)
		first := replace(digits[0])
		last := replace(digits[len(digits)-1])
		first_last := fmt.Sprintf("%s%s", first, last)
		val, err := strconv.Atoi(first_last)
		if err != nil {
			panic(err)
		}
		sum += val
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	firstRegex := regexp.MustCompile("(" + regex + ")")
	lastRegex := regexp.MustCompile(".*" + firstRegex.String())
	sum := 0
	for _, line := range lines {
		first := firstRegex.FindStringSubmatch(line)[1]
		last := lastRegex.FindStringSubmatch(line)[1]
		first_last := fmt.Sprintf("%s%s", replace(first), replace(last))
		val, err := strconv.Atoi(first_last)
		if err != nil {
			panic(err)
		}
		sum += val
	}
	fmt.Println(sum)
}

func replace(line string) string {
	temp := line
	for i, digit := range digits {
		temp = strings.ReplaceAll(temp, digit, fmt.Sprint(i+1))
	}
	return temp
}
