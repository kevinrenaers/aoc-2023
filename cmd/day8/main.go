package main

import (
	"advent-of-code/internal"
	"fmt"
	"regexp"
	"strings"
)

var node_regex = "([A-Z]+) = \\(([A-Z]+), ([A-Z]+)\\)"

func main() {
	lines, err := internal.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	leftRightList := strings.Split(lines[0], "")
	nodeMap := make(map[string]node)
	node_re := regexp.MustCompile(node_regex)
	for i := 2; i < len(lines); i++ {
		lineGroups := node_re.FindStringSubmatch(lines[i])
		nodeMap[lineGroups[1]] = node{current: lineGroups[1], left: lineGroups[2], right: lineGroups[3]}
	}
	currentPos := "AAA"
	var steps int
	for currentPos != "ZZZ" {
		lr := leftRightList[steps%len(leftRightList)]
		currentPos = nodeMap[currentPos].getNext(lr)
		steps += 1
	}
	fmt.Println(steps)
}

func part2(lines []string) {
	leftRightList := strings.Split(lines[0], "")
	nodeMap := make(map[string]node)
	node_re := regexp.MustCompile(node_regex)
	for i := 2; i < len(lines); i++ {
		lineGroups := node_re.FindStringSubmatch(lines[i])
		nodeMap[lineGroups[1]] = node{current: lineGroups[1], left: lineGroups[2], right: lineGroups[3]}
	}
	var startingPositions []string
	for key := range nodeMap {
		if strings.HasSuffix(key, "A") {
			startingPositions = append(startingPositions, key)
		}
	}
	endingPositions := make([]int, len(startingPositions))
	for i, startPos := range startingPositions {
		curPos := startPos
		endingFound := false
		var steps int
		for !endingFound {
			lr := leftRightList[steps%len(leftRightList)]
			curPos = nodeMap[curPos].getNext(lr)
			if strings.HasSuffix(curPos, "Z") {
				endingPositions[i] = steps + 1
				endingFound = true
			}
			steps += 1
		}
	}
	lcm := endingPositions[0]
	for i, endPos := range endingPositions {
		if i == 0 {
			lcm = endPos
			continue
		}
		lcm = LCM(lcm, endPos)
	}
	fmt.Println(lcm)
}

type node struct {
	current string
	left    string
	right   string
}

func (n node) getNext(lr string) string {
	if lr == "L" {
		return n.left
	}
	return n.right
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
