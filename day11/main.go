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
		fmt.Println("part 1:", part1(input, 2))
	}

	if *p2 {
		fmt.Println("part 2:", part2(input, 1000000))
	}
}

func part1(input string, gap int) int {
	return solve(input, gap)
}

func part2(input string, gap int) int {
	return solve(input, gap)
}

func solve(input string, gap int) int {
	grid := parse(input)
	m := len(grid)
	n := len(grid[0])
	preSumRow := []int{0}
	preSumCol := []int{0}

	for i := 0; i < m; i++ {
		preSumRow = append(preSumRow, preSumRow[len(preSumRow)-1])
		if !strings.Contains(grid[i], "#") {
			preSumRow[len(preSumRow)-1] += 1
		}
	}

	for j := 0; j < n; j++ {
		hasGalaxy := false
		for i := 0; i < m; i++ {
			if grid[i][j] == '#' {
				hasGalaxy = true
				break
			}
		}

		preSumCol = append(preSumCol, preSumCol[len(preSumCol)-1])
		if !hasGalaxy {
			preSumCol[len(preSumCol)-1] += 1
		}
	}

	res := 0
	var galaxies []cell
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '#' {
				galaxies = append(galaxies, cell{i, j})
			}
		}
	}

	for i := 0; i < len(galaxies); i++ {
		r1 := galaxies[i].row + preSumRow[galaxies[i].row+1]*(gap-1)
		c1 := galaxies[i].col + preSumCol[galaxies[i].col+1]*(gap-1)

		for j := i + 1; j < len(galaxies); j++ {
			r2 := galaxies[j].row + preSumRow[galaxies[j].row+1]*(gap-1)
			c2 := galaxies[j].col + preSumCol[galaxies[j].col+1]*(gap-1)
			res += abs(r1-r2) + abs(c1-c2)
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
