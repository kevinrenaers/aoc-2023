package main

import (
	"advent-of-code/internal"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var input_regex = "[a-zA-Z]+:"
var space_regex = " {2,}"

func main() {
	lines, err := internal.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	input_re := regexp.MustCompile(input_regex)
	space_re := regexp.MustCompile(space_regex)
	inputLine := input_re.ReplaceAllString(lines[0], "")
	inputLine = strings.TrimPrefix(space_re.ReplaceAllString(inputLine, " "), " ")
	timeStrings := strings.Split(inputLine, " ")
	distanceLine := input_re.ReplaceAllString(lines[1], "")
	distanceLine = strings.TrimPrefix(space_re.ReplaceAllString(distanceLine, " "), " ")
	distanceStrings := strings.Split(distanceLine, " ")
	product := 1
	for i, timeString := range timeStrings {
		var amountWinning int
		time, _ := strconv.Atoi(timeString)
		distance, _ := strconv.Atoi(distanceStrings[i])
		for j := 1; j < time; j++ {
			calculatedDistance := j * (time - j)
			if calculatedDistance > distance {
				amountWinning += 1
			}
		}
		product *= amountWinning
	}
	fmt.Println(product)
}

func part2(lines []string) {
	input_re := regexp.MustCompile(input_regex)
	space_re := regexp.MustCompile(space_regex)
	inputLine := input_re.ReplaceAllString(lines[0], "")
	inputLine = space_re.ReplaceAllString(inputLine, "")
	distanceLine := input_re.ReplaceAllString(lines[1], "")
	distanceLine = space_re.ReplaceAllString(distanceLine, "")
	var amountWinning int
	time, _ := strconv.Atoi(inputLine)
	distance, _ := strconv.Atoi(distanceLine)
	for j := 1; j < time; j++ {
		calculatedDistance := j * (time - j)
		if calculatedDistance > distance {
			amountWinning += 1
		}
	}
	fmt.Println(amountWinning)
}
