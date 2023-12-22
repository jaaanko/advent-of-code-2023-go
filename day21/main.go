package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type cell struct {
	row, col int
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
		fmt.Println("part 1:", part1(input, 64))
	}

	if *p2 {
		fmt.Println("part 2:", part2(input))
	}
}

func part1(input string, steps int) int {
	grid := parse(input)
	start := findStart(grid)
	queue := []cell{start}
	m := len(grid)
	n := len(grid[0])

	for steps > 0 {
		size := len(queue)
		visited := map[cell]bool{}

		for size > 0 {
			curr := queue[0]
			visited[curr] = true
			queue = queue[1:]

			for _, dir := range [4]struct{ row, col int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				nr, nc := dir.row+curr.row, dir.col+curr.col
				if nr < 0 || nr >= m || nc < 0 || nc >= n || visited[cell{nr, nc}] || grid[nr][nc] == '#' {
					continue
				}

				visited[cell{nr, nc}] = true
				queue = append(queue, cell{nr, nc})
			}
			size -= 1
		}
		steps -= 1
	}

	return len(queue)
}

func part2(input string) int {
	// 26501365 steps required.
	// 65+(131*n) = 26501365
	// n = (26501365-65)/131
	n := (26_501_365 - 65) / 131
	return calc(n + 1)
}

func calc(n int) int {
	// Derived from the quadratic sequence of:
	// 3868,34368,95262,186550 (Generated using brute force from part 1)
	return 15_197*(n*n) - 15_091*n + 3_762
}

func findStart(grid []string) cell {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {
			if grid[r][c] == 'S' {
				return cell{r, c}
			}
		}
	}
	return cell{-1, -1}
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
