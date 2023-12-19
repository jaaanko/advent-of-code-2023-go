package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type instruction struct {
	dir   string
	dist  int
	color string
}

type delta struct {
	row, col int
}

var deltaLookup = map[string]delta{
	"U": {-1, 0},
	"D": {1, 0},
	"L": {0, -1},
	"R": {0, 1},
}

type block struct {
	r1, c1, r2, c2 int
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
	instructions := parse(input)
	return solve(instructions)
}

func part2(input string) int {
	instructions := parse(input)
	mapping := map[byte]string{
		'0': "R",
		'1': "D",
		'2': "L",
		'3': "U",
	}

	for i := range instructions {
		n := len(instructions[i].color)
		hex := instructions[i].color[1 : n-1]
		decimalInt, _ := strconv.ParseInt(hex, 16, 64)
		instructions[i].dir = mapping[instructions[i].color[n-1]]
		instructions[i].dist = int(decimalInt)
	}

	return solve(instructions)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func solve(instructions []instruction) int {
	rows := []int{0, 1}
	cols := []int{0, 1}
	seenRows := map[int]bool{0: true, 1: true}
	seenCols := map[int]bool{0: true, 1: true}
	row := 0
	col := 0
	var blocks []block

	for _, ins := range instructions {
		r1 := min(row, row+ins.dist*deltaLookup[ins.dir].row)
		r2 := max(row, row+ins.dist*deltaLookup[ins.dir].row)
		c1 := min(col, col+ins.dist*deltaLookup[ins.dir].col)
		c2 := max(col, col+ins.dist*deltaLookup[ins.dir].col)
		row += ins.dist * deltaLookup[ins.dir].row
		col += ins.dist * deltaLookup[ins.dir].col
		r2 += 1
		c2 += 1
		blocks = append(blocks, block{r1, c1, r2, c2})

		if !seenRows[r1] {
			seenRows[r1] = true
			rows = append(rows, r1)
		}
		if !seenRows[r2] {
			seenRows[r2] = true
			rows = append(rows, r2)
		}

		if !seenCols[c1] {
			seenCols[c1] = true
			cols = append(cols, c1)
		}
		if !seenCols[c2] {
			seenCols[c2] = true
			cols = append(cols, c2)
		}
	}

	sort.Ints(rows)
	sort.Ints(cols)
	uncompressRow := map[int]int{}
	uncompressCol := map[int]int{}
	compressRow := map[int]int{}
	compressCol := map[int]int{}
	row = 0
	col = 0

	for i := 0; i < len(rows); i++ {
		compressRow[rows[i]] = row
		uncompressRow[row] = rows[i]
		row += 1
	}

	for i := 0; i < len(cols); i++ {
		compressCol[cols[i]] = col
		uncompressCol[col] = cols[i]
		col += 1
	}

	m := row
	n := col
	var grid [][]rune
	for r := 0; r < m; r++ {
		var row []rune
		for c := 0; c < n; c++ {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}

	for _, block := range blocks {
		for r := compressRow[block.r1]; r < compressRow[block.r2]; r++ {
			for c := compressCol[block.c1]; c < compressCol[block.c2]; c++ {
				grid[r][c] = '#'
			}
		}
	}

	color := 'x'
	for r := 0; r < m; r++ {
		paint(grid, r, 0, color)
		paint(grid, r, n-1, color)
	}

	for c := 0; c < n; c++ {
		paint(grid, 0, c, color)
		paint(grid, m-1, c, color)
	}

	res := 0
	for r := 0; r < m-1; r++ {
		for c := 0; c < n-1; c++ {
			if grid[r][c] == color {
				continue
			}
			width := uncompressCol[c+1] - uncompressCol[c]
			height := uncompressRow[r+1] - uncompressRow[r]
			res += height * width
		}
	}

	return res
}

func paint(grid [][]rune, r, c int, color rune) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '#' || grid[r][c] == color {
		return
	}
	grid[r][c] = color
	for _, v := range deltaLookup {
		paint(grid, r+v.row, c+v.col, color)
	}
}

func parse(input string) []instruction {
	var instructions []instruction
	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		ins := strings.Fields(line)
		dist, _ := strconv.Atoi(ins[1])
		color := ins[2][1 : len(ins[2])-1]
		instructions = append(instructions, instruction{dir: ins[0], dist: dist, color: color})
	}

	return instructions
}
