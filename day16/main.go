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

type direction struct {
	row, col int
}

var (
	left  = direction{row: 0, col: -1}
	right = direction{row: 0, col: 1}
	up    = direction{row: -1, col: 0}
	down  = direction{row: 1, col: 0}
)

type state struct {
	cell cell
	dir  direction
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
	return countEnergized(grid, state{cell: cell{0, 0}, dir: right})

	// var tmp [][]rune

	// for i := 0; i < len(grid); i++ {
	// 	tmp = append(tmp, []rune(grid[i]))
	// }

	// for k := range res {
	// 	tmp[k.row][k.col] = '#'
	// }

	// for i := 0; i < len(grid); i++ {
	// 	fmt.Println(string(tmp[i]))
	// }

}

func countEnergized(grid []string, start state) int {
	visited := map[state]bool{start: true}
	queue := []state{start}
	res := map[cell]bool{start.cell: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nei := range getNeighbours(grid, curr) {
			if !visited[nei] {
				visited[nei] = true
				res[nei.cell] = true
				queue = append(queue, nei)
			}
		}
	}

	return len(res)
}

func part2(input string) int {
	grid := parse(input)
	res := 0
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		res = max(
			res,
			countEnergized(grid, state{cell: cell{i, 0}, dir: right}),
			countEnergized(grid, state{cell: cell{i, n - 1}, dir: left}),
		)
	}

	for j := 0; j < n; j++ {
		res = max(
			res,
			countEnergized(grid, state{cell: cell{0, j}, dir: down}),
			countEnergized(grid, state{cell: cell{m - 1, j}, dir: up}),
		)
	}

	return res
}

func getNeighbours(grid []string, curr state) []state {
	r := curr.cell.row
	c := curr.cell.col
	m := len(grid)
	n := len(grid[0])
	nr := curr.cell.row + curr.dir.row
	nc := curr.cell.col + curr.dir.col

	if curr.dir == left || curr.dir == right {
		if grid[r][c] == '-' || grid[r][c] == '.' {
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				return nil
			}
			return []state{{cell: cell{nr, nc}, dir: curr.dir}}
		}
		if grid[r][c] == '|' {
			var neighbours []state
			for _, dir := range []direction{up, down} {
				if r+dir.row >= 0 && r+dir.row < m {
					neighbours = append(neighbours, state{cell: cell{r + dir.row, c}, dir: dir})
				}
			}
			return neighbours
		}

		if r+down.row < m && (grid[r][c] == '/' && curr.dir == left || grid[r][c] == '\\' && curr.dir == right) {
			return []state{{cell: cell{r + down.row, c}, dir: down}}
		}
		if r+up.row >= 0 && (grid[r][c] == '/' && curr.dir == right || grid[r][c] == '\\' && curr.dir == left) {
			return []state{{cell: cell{r + up.row, c}, dir: up}}
		}
		return nil
	}

	if grid[r][c] == '|' || grid[r][c] == '.' {
		if nr < 0 || nr >= m || nc < 0 || nc >= n {
			return nil
		}
		return []state{{cell: cell{nr, nc}, dir: curr.dir}}
	}

	if grid[r][c] == '-' {
		var neighbours []state
		for _, dir := range []direction{left, right} {
			if c+dir.col >= 0 && c+dir.col < n {
				neighbours = append(neighbours, state{cell: cell{r, c + dir.col}, dir: dir})
			}
		}
		return neighbours
	}

	if c+right.col < n && (grid[r][c] == '/' && curr.dir == up || grid[r][c] == '\\' && curr.dir == down) {
		return []state{{cell: cell{r, c + right.col}, dir: right}}
	}
	if c+left.col >= 0 && (grid[r][c] == '/' && curr.dir == down || grid[r][c] == '\\' && curr.dir == up) {
		return []state{{cell: cell{r, c + left.col}, dir: left}}
	}

	return nil
}

func max(nums ...int) int {
	res := math.MinInt
	for _, num := range nums {
		if num > res {
			res = num
		}
	}
	return res
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
