package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

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
	moves, adj := parse(input)
	curr := "AAA"
	target := "ZZZ"
	n := len(moves)
	i := 0
	res := 0

	for curr != target {
		if moves[i%n] == 'L' {
			curr = adj[curr][0]
		} else {
			curr = adj[curr][1]
		}
		i += 1
		res += 1
	}

	return res
}

func part2(input string) int {
	moves, adj := parse(input)
	var starts []string

	for k := range adj {
		if k[len(k)-1] == 'A' {
			starts = append(starts, k)
		}
	}

	var numMoves []int
	for _, start := range starts {
		i, err := firstEndsWithZ(start, moves, adj)
		if err != nil {
			panic(err)
		}
		numMoves = append(numMoves, i)
	}

	return lcm(numMoves...)
}

func firstEndsWithZ(curr, moves string, adj map[string][]string) (int, error) {
	type entry struct {
		node    string
		moveIdx int
	}

	seen := map[entry]bool{}
	n := len(moves)
	i := 0

	for !seen[entry{curr, i % n}] {
		seen[entry{curr, i % n}] = true
		if curr[len(curr)-1] == 'Z' {
			return i, nil
		}

		if moves[i%n] == 'L' {
			curr = adj[curr][0]
		} else {
			curr = adj[curr][1]
		}
		i += 1
	}

	return -1, errors.New("no valid path found")
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func lcm(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res = (res * nums[i]) / gcd(res, nums[i])
	}

	return res
}

func parse(input string) (string, map[string][]string) {
	blocks := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")
	adj := map[string][]string{}
	moves := blocks[0]

	for _, line := range strings.Split(blocks[1], "\n") {
		contents := strings.Split(line, " = ")
		source := contents[0]
		dest := strings.Split(contents[1], ", ")
		left := dest[0][1:len(dest[0])]
		right := dest[1][0 : len(dest[1])-1]

		adj[source] = append(adj[source], left)
		adj[source] = append(adj[source], right)
	}

	return moves, adj
}
