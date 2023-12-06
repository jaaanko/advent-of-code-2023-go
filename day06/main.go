package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time, dist int
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
	races := parse(input)
	res := 1

	for _, race := range races {
		ways := 0
		for i := 0; i < race.time; i++ {
			remaining := race.time - i
			traveled := remaining * i
			if traveled > race.dist {
				ways += 1
			}
		}
		res *= ways
	}

	return res
}

func part2(input string) int {
	races := parse(input)
	var timeStr, distStr string

	for _, race := range races {
		timeStr += strconv.Itoa(race.time)
		distStr += strconv.Itoa(race.dist)
	}

	time, _ := strconv.Atoi(timeStr)
	dist, _ := strconv.Atoi(distStr)
	ways := 0

	for i := 0; i < time; i++ {
		remaining := time - i
		traveled := remaining * i
		if traveled > dist {
			ways += 1
		}
	}

	return ways
}

func parse(input string) []race {
	lines := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
	times := strings.Fields(lines[0])
	dists := strings.Fields(lines[1])
	var races []race

	for i := 1; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		dist, _ := strconv.Atoi(dists[i])
		races = append(races, race{time: time, dist: dist})
	}

	return races
}
