package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type coord struct {
	x, y int
}

type number struct {
	val             int
	adjSymbolCoords map[coord]bool
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
	grid := parse(input)
	nums := findNums(grid)
	res := 0

	for _, num := range nums {
		if len(num.adjSymbolCoords) > 0 {
			res += num.val
		}
	}

	return res
}

func part2(input string) int {
	grid := parse(input)
	nums := findNums(grid)
	adjNumbers := make(map[coord][]int)
	res := 0

	for _, num := range nums {
		for coord := range num.adjSymbolCoords {
			if grid[coord.y][coord.x] == '*' {
				adjNumbers[coord] = append(adjNumbers[coord], num.val)
			}
		}
	}

	for coord, nums := range adjNumbers {
		if grid[coord.y][coord.x] == '*' && len(nums) == 2 {
			res += nums[0] * nums[1]
		}
	}

	return res
}

func parse(input string) [][]rune {
	var grid [][]rune

	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}

func findNums(grid [][]rune) []number {
	var res []number
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		curr := 0
		adjSymbolCoords := make(map[coord]bool)

		for j := 0; j < n; j++ {
			if unicode.IsDigit(grid[i][j]) {
				curr = curr*10 + int(grid[i][j]-'0')
				for di := -1; di < 2; di++ {
					for dj := -1; dj < 2; dj++ {
						if di == 0 && dj == 0 {
							continue
						}

						ni, nj := i+di, j+dj
						if 0 <= ni && ni < m && 0 <= nj && nj < n && !unicode.IsDigit(grid[ni][nj]) && grid[ni][nj] != '.' {
							adjSymbolCoords[coord{y: ni, x: nj}] = true
						}

					}
				}

				if j == n-1 || !unicode.IsDigit(grid[i][j+1]) {
					res = append(res, number{val: curr, adjSymbolCoords: adjSymbolCoords})
					curr = 0
					adjSymbolCoords = make(map[coord]bool)
				}
			}
		}
	}

	return res
}
