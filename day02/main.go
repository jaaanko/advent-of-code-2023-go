package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type bag struct {
	colorCount map[string]int
}

type game struct {
	id   int
	bags []bag
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
	have := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	games := parse(input)
	res := 0

	for _, game := range games {
		valid := true
		for _, bag := range game.bags {
			for color, count := range bag.colorCount {
				if count > have[color] {
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			res += game.id
		}
	}

	return res
}

func part2(input string) int {
	games := parse(input)
	res := 0
	colors := []string{"red", "green", "blue"}

	for _, game := range games {
		maxCount := make(map[string]int)
		for _, bag := range game.bags {
			for _, color := range colors {
				if bag.colorCount[color] > maxCount[color] {
					maxCount[color] = bag.colorCount[color]
				}
			}
		}
		prod := 1
		for _, color := range colors {
			prod *= maxCount[color]
		}
		res += prod
	}

	return res
}

func parse(input string) []game {
	var games []game
	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		var gameId int
		line := strings.Split(line, ": ")
		fmt.Sscanf(line[0], "Game %d", &gameId)
		game := game{id: gameId}

		for _, contents := range strings.Split(line[1], "; ") {
			bag := bag{colorCount: make(map[string]int)}
			for _, content := range strings.Split(contents, ", ") {
				var count int
				var color string
				fmt.Sscanf(content, "%d %s", &count, &color)
				bag.colorCount[color] = count
			}
			game.bags = append(game.bags, bag)
		}
		games = append(games, game)
	}

	return games
}
