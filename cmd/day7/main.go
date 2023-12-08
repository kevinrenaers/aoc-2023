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
	var hands []hand
	for _, line := range lines {
		input := strings.Split(line, " ")
		cards := strings.Split(input[0], "")
		bid, _ := strconv.Atoi(input[1])
		hands = append(hands, newHand(cards, bid, false))
	}
	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]
		if a.handType > b.handType {
			return true
		} else if a.handType < b.handType {
			return false
		}
		for k, card := range a.cards {
			bCard := b.cards[k]
			if card == b.cards[k] {
				continue
			}
			if cardMapping[card] > cardMapping[bCard] {
				return true
			} else {
				return false
			}
		}
		return false
	})
	var sum int
	for i, hand := range hands {
		sum += hand.bid * (len(hands) - i)
	}
	fmt.Println(sum)
}

func part2(lines []string) {
	var hands []hand
	for _, line := range lines {
		input := strings.Split(line, " ")
		cards := strings.Split(input[0], "")
		bid, _ := strconv.Atoi(input[1])
		hands = append(hands, newHand(cards, bid, true))
	}
	sort.Slice(hands, func(i, j int) bool {
		a := hands[i]
		b := hands[j]
		if a.handType > b.handType {
			return true
		} else if a.handType < b.handType {
			return false
		}
		for k, card := range a.cards {
			bCard := b.cards[k]
			if card == b.cards[k] {
				continue
			}
			if cardMappingPartTwo[card] > cardMappingPartTwo[bCard] {
				return true
			} else {
				return false
			}
		}
		return false
	})
	var sum int
	for i, hand := range hands {
		sum += hand.bid * (len(hands) - i)
	}
	fmt.Println(sum)
}

var cardMapping = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": 9,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

var cardMappingPartTwo = map[string]int{
	"A": 12,
	"K": 11,
	"Q": 10,
	"J": -1,
	"T": 8,
	"9": 7,
	"8": 6,
	"7": 5,
	"6": 4,
	"5": 3,
	"4": 2,
	"3": 1,
	"2": 0,
}

type hand_type int

var (
	FiveOAKind  = hand_type(6)
	FourOAKind  = hand_type(5)
	FullHouse   = hand_type(4)
	ThreeOAKind = hand_type(3)
	TwoPair     = hand_type(2)
	OnePair     = hand_type(1)
	HighCard    = hand_type(0)
)

type hand struct {
	cards    []string
	bid      int
	handType hand_type
}

func newHand(cards []string, bid int, useJoker bool) hand {
	out := hand{
		cards:    cards,
		bid:      bid,
		handType: calc_type(cards, useJoker),
	}
	return out
}

func calc_type(cards []string, useJoker bool) hand_type {
	cardMap := make(map[string]int)
	var jokerCount int
	for _, card := range cards {
		if card == "J" {
			jokerCount += 1
		}
		cardMap[card] += 1
	}

	if jokerCount == 5 {
		return FiveOAKind
	}
	var highestCountCard string
	var highestCount int
	if useJoker {
		for card, count := range cardMap {
			if card == "J" {
				continue
			}
			if count > highestCount {
				highestCount = count
				highestCountCard = card
			}
		}
	}
	if useJoker {
		cardMap[highestCountCard] += jokerCount
		delete(cardMap, "J")
	}
	threeMatching := false
	pairs := 0
	for _, count := range cardMap {
		switch count {
		case 5:
			return FiveOAKind
		case 4:
			return FourOAKind
		case 3:
			threeMatching = true
			break
		case 2:
			pairs += 1
			break
		}
	}
	if threeMatching {
		if pairs == 1 {
			return FullHouse
		}
		return ThreeOAKind
	}
	if pairs == 2 {
		return TwoPair
	}
	if pairs == 1 {
		return OnePair
	}
	return HighCard
}
