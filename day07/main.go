package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

var cards = [...]rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

type handType int

const (
	highCard handType = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

type hand struct {
	cards    [5]rune
	bid      int
	handType handType
}

func main() {
	p1 := flag.Bool("p1", false, "run part 1")
	p2 := flag.Bool("p2", false, "run part 2")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatal("no input file provided")
	}

	b, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}

	input := string(b)

	if *p1 {
		fmt.Println("part 1:", part1(input))
	}

	if *p2 {
		fmt.Println("part 2:", part2(input))
	}
}

func part1(input string) int {
	hands := parse(input, false)
	strength := map[rune]int{}

	for i, card := range cards {
		strength[card] = len(cards) - i
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for k := 0; k < 5; k++ {
				if strength[hands[i].cards[k]] == strength[hands[j].cards[k]] {
					continue
				}
				return strength[hands[i].cards[k]] < strength[hands[j].cards[k]]
			}
			return false
		}
		return hands[i].handType < hands[j].handType
	})

	res := 0
	for i, hand := range hands {
		res += hand.bid * (i + 1)
	}

	return res
}

func part2(input string) int {
	hands := parse(input, true)
	strength := map[rune]int{}

	for i, card := range cards {
		strength[card] = len(cards) - i
	}
	strength['J'] = -1

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for k := 0; k < 5; k++ {
				if strength[hands[i].cards[k]] == strength[hands[j].cards[k]] {
					continue
				}
				return strength[hands[i].cards[k]] < strength[hands[j].cards[k]]
			}
			return false
		}
		return hands[i].handType < hands[j].handType
	})

	res := 0
	for i, hand := range hands {
		res += hand.bid * (i + 1)
	}

	return res
}

func parse(input string, wildcardEnabled bool) []hand {
	var cards string
	var bid int
	var hands []hand

	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		fmt.Sscanf(line, "%s %d", &cards, &bid)
		hands = append(hands, newHand([5]rune([]rune(cards)), bid, wildcardEnabled))
	}

	return hands
}

func newHand(cards [5]rune, bid int, wildcardEnabled bool) hand {
	return hand{
		cards:    cards,
		bid:      bid,
		handType: getHandType(cards, wildcardEnabled),
	}
}

func getHandType(cards [5]rune, wildcardEnabled bool) handType {
	counts := map[rune]int{}
	wildcards := 0
	for _, card := range cards {
		if wildcardEnabled && card == 'J' {
			wildcards += 1
		} else {
			counts[card] += 1
		}
	}

	var values []int
	for _, v := range counts {
		values = append(values, v)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i] > values[j]
	})

	if wildcardEnabled {
		if len(values) == 0 {
			return fiveOfAKind
		}
		values[0] += wildcards
	}

	if values[0] == 5 {
		return fiveOfAKind
	}
	if values[0] == 4 {
		return fourOfAKind
	}
	if values[0] == 3 {
		if values[1] == 2 {
			return fullHouse
		}
		return threeOfAKind
	}
	if values[0] == 2 {
		if values[1] == 2 {
			return twoPair
		}
		return onePair
	}

	return highCard
}
