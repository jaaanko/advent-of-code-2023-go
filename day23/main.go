package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type cell struct {
	row, col int
}

type delta struct {
	row, col int
}

var (
	leftDelta  = delta{0, -1}
	rightDelta = delta{0, 1}
	upDelta    = delta{-1, 0}
	downDelta  = delta{1, 0}
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
	grid := parse(input)
	return recurse(grid, cell{0, 1}, map[cell]bool{}, false) - 1
}

func part2(input string) int {
	grid := parse(input)
	res := recurse(grid, cell{0, 1}, map[cell]bool{}, true) - 1
	return res
}

func recurse(grid []string, c cell, visited map[cell]bool, ignoreSlopes bool) int {
	if c.row == len(grid)-1 && c.col == len(grid[0])-2 {
		return 1
	}

	res := math.MinInt
	visited[c] = true
	for _, d := range [4]delta{upDelta, downDelta, leftDelta, rightDelta} {
		nr, nc := d.row+c.row, d.col+c.col

		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) || grid[nr][nc] == '#' || visited[cell{nr, nc}] {
			continue
		}

		if !ignoreSlopes {
			if grid[c.row][c.col] == '>' && d != rightDelta || grid[c.row][c.col] == 'v' && d != downDelta {
				continue
			}

			if d == upDelta && grid[nr][nc] == 'v' {
				continue
			}

			if d == leftDelta && grid[nr][nc] == '>' {
				continue
			}
		}

		res = max(res, recurse(grid, cell{nr, nc}, visited, ignoreSlopes)+1)
	}
	visited[c] = false
	return res
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
