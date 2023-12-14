package main

import (
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
		fmt.Println("part 2:", part2(input, 1000000000))
	}
}

const (
	north = 1
	west  = 1
	south = -1
	east  = -1
)

func part1(input string) int {
	grid := parse(input)
	return calcLoad(tiltVertical(grid, 1))
}

func calcLoad(grid []string) int {
	m := len(grid)
	n := len(grid[0])
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 'O' {
				res += m - i
			}
		}
	}

	return res
}

func part2(input string, spins int) int {
	grid := parse(input)
	seen := map[string]int{}
	curr := getStringRep(grid)
	cycle := 0

	for {
		if _, ok := seen[curr]; ok {
			break
		}
		seen[curr] = cycle
		cycle += 1
		grid = spin(grid)
		curr = getStringRep(grid)
	}

	delta := cycle - seen[curr]
	offset := seen[curr] - 1
	remaining := (spins-offset)%delta - 1

	for remaining > 0 {
		grid = spin(grid)
		remaining -= 1
	}

	return calcLoad(grid)
}

func spin(grid []string) []string {
	rotated := tiltVertical(grid, north)
	rotated = tiltHorizontal(rotated, west)
	rotated = tiltVertical(rotated, south)
	rotated = tiltHorizontal(rotated, east)

	return rotated
}

func tiltVertical(grid []string, di int) []string {
	var tilted [][]rune
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		var row []rune
		for j := 0; j < n; j++ {
			row = append(row, rune(grid[i][j]))
		}
		tilted = append(tilted, row)
	}

	var start int
	if di == south {
		start = m - 1
	}

	for j := 0; j < n; j++ {
		var prev int
		if di == south {
			prev = m
		} else {
			prev = -1
		}

		for i := start; i >= 0 && i < m; i += di {
			if grid[i][j] == '#' {
				prev = i
			} else if grid[i][j] == 'O' {
				prev += di
				tilted[i][j] = '.'
				tilted[prev][j] = 'O'
			}
		}
	}

	var res []string
	for i := 0; i < m; i++ {
		res = append(res, string(tilted[i]))
	}

	return res
}

func tiltHorizontal(grid []string, dj int) []string {
	var tilted [][]rune
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		var row []rune
		for j := 0; j < n; j++ {
			row = append(row, rune(grid[i][j]))
		}
		tilted = append(tilted, row)
	}

	var start int
	if dj == east {
		start = n - 1
	}

	for i := 0; i < m; i++ {
		var prev int
		if dj == east {
			prev = n
		} else {
			prev = -1
		}

		for j := start; j >= 0 && j < n; j += dj {
			if grid[i][j] == '#' {
				prev = j
			} else if grid[i][j] == 'O' {
				prev += dj
				tilted[i][j] = '.'
				tilted[i][prev] = 'O'
			}
		}
	}

	var res []string
	for i := 0; i < m; i++ {
		res = append(res, string(tilted[i]))
	}

	return res
}

func getStringRep(grid []string) string {
	var sb strings.Builder
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			sb.WriteByte(grid[i][j])
		}
		sb.WriteByte('\n')
	}

	return sb.String()
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
