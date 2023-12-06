package main

import (
	"advent-of-code/internal"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var game_regex = "Game ([0-9]+): "

func main() {
	lines, err := internal.ReadLines("input.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	game_re := regexp.MustCompile(game_regex)
	var err error
	sumPossible := 0
	for _, line := range lines {
		game := Game{}
		game.id, err = strconv.Atoi(game_re.FindStringSubmatch(line)[1])
		if err != nil {
			panic(err)
		}
		line = game_re.ReplaceAllString(line, "")
		game = game.ParseInput(line)
		if game.Check(12, 14, 13) {
			sumPossible += game.id
		}
	}
	fmt.Println(sumPossible)
}

func part2(lines []string) {
	game_re := regexp.MustCompile(game_regex)
	var err error
	sumPower := 0
	for _, line := range lines {
		game := Game{}
		game.id, err = strconv.Atoi(game_re.FindStringSubmatch(line)[1])
		if err != nil {
			panic(err)
		}
		line = game_re.ReplaceAllString(line, "")
		game = game.ParseInput(line)
		game = game.CountNeeded()
		sumPower += (game.greenNeeded * game.redNeeded * game.blueNeeded)
	}
	fmt.Println(sumPower)
}

type Game struct {
	id          int
	greenNeeded int
	redNeeded   int
	blueNeeded  int
	greenDices  []int
	redDices    []int
	blueDices   []int
}

func (g Game) ParseInput(input string) Game {
	countRegex := regexp.MustCompile("[0-9]+")
	colorRegex := regexp.MustCompile("red|blue|green")
	showings := strings.Split(input, "; ")
	for _, showing := range showings {
		dices := strings.Split(showing, ", ")
		for _, diceCount := range dices {
			count, err := strconv.Atoi(countRegex.FindString(diceCount))
			if err != nil {
				panic(err)
			}
			switch colorRegex.FindString(diceCount) {
			case "red":
				g.redDices = append(g.redDices, count)
			case "blue":
				g.blueDices = append(g.blueDices, count)
			case "green":
				g.greenDices = append(g.greenDices, count)
			}
		}
	}
	return g
}

func (g Game) Check(redCount, blueCount, greenCount int) bool {
	for _, count := range g.redDices {
		if count > redCount {
			return false
		}
	}
	for _, count := range g.blueDices {
		if count > blueCount {
			return false
		}
	}
	for _, count := range g.greenDices {
		if count > greenCount {
			return false
		}
	}
	return true
}

func (g Game) CountNeeded() Game {
	for _, count := range g.redDices {
		if count > g.redNeeded {
			g.redNeeded = count
		}
	}
	for _, count := range g.blueDices {
		if count > g.blueNeeded {
			g.blueNeeded = count
		}
	}
	for _, count := range g.greenDices {
		if count > g.greenNeeded {
			g.greenNeeded = count
		}
	}
	return g
}
