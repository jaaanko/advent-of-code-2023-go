package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type card struct {
	winningNums map[int]bool
	ownedNums   []int
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
	cards := parse(input)
	res := 0

	for _, card := range cards {
		wins := countWins(card)
		if wins > 0 {
			res += 1 << (wins - 1)
		}
	}

	return res
}

func part2(input string) int {
	cards := parse(input)
	count := map[int]int{}
	res := len(cards)

	for i, card := range cards {
		count[i] += 1
		for j := i + 1; j <= i+countWins(card); j++ {
			count[j] += count[i]
			res += count[i]
		}
	}

	return res
}

func parse(input string) []card {
	var cards []card
	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		line := strings.Split(line, ": ")
		nums := strings.Split(line[1], " | ")
		card := card{winningNums: map[int]bool{}}

		for _, num := range strings.Fields(nums[0]) {
			num, _ := strconv.Atoi(num)
			card.winningNums[num] = true
		}
		for _, num := range strings.Fields(nums[1]) {
			num, _ := strconv.Atoi(num)
			card.ownedNums = append(card.ownedNums, num)
		}

		cards = append(cards, card)
	}

	return cards
}

func countWins(c card) int {
	count := 0
	for _, num := range c.ownedNums {
		if c.winningNums[num] {
			count += 1
		}
	}

	return count
}
