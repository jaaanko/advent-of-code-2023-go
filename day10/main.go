package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type node struct {
	row, col int
}

type direction int

const (
	north direction = iota
	south
	east
	west
)

type delta struct {
	row, col int
}

var directionLookup = map[delta]direction{
	{-1, 0}: north,
	{1, 0}:  south,
	{0, 1}:  east,
	{0, -1}: west,
}

var reverseLookup = map[direction]direction{
	north: south,
	south: north,
	east:  west,
	west:  east,
}

var mapping = map[direction]map[byte]bool{
	north: {'|': true, 'L': true, 'J': true},
	south: {'|': true, '7': true, 'F': true},
	east:  {'-': true, 'L': true, 'F': true},
	west:  {'-': true, '7': true, 'J': true},
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
	start := findStart(grid)
	adj := buildGraph(start, grid)
	dist := map[node]int{start: 0}
	queue := []node{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, nei := range adj[curr] {
			if _, ok := dist[nei]; ok {
				continue
			}
			dist[nei] = dist[curr] + 1
			queue = append(queue, nei)
		}
	}

	res := 0
	for _, v := range dist {
		if v > res {
			res = v
		}
	}

	return res
}

func part2(input string) int {
	grid := parse(input)
	m := len(grid)
	n := len(grid[0])
	start := findStart(grid)
	adj := buildGraph(start, grid)
	var cleanedGrid []string

	for i := 0; i < m; i++ {
		var row strings.Builder
		for j := 0; j < n; j++ {
			if _, ok := adj[node{i, j}]; !ok {
				row.WriteByte('*')
			} else if start.row == i && start.col == j {
				row.WriteByte('S')
			} else {
				row.WriteByte(grid[i][j])
			}
		}
		cleanedGrid = append(cleanedGrid, row.String())
	}

	expandedGrid := expandGrid(cleanedGrid)
	m = len(expandedGrid)
	n = len(expandedGrid[0])
	colored := map[node]bool{}

	for j := 0; j < n; j++ {
		if (expandedGrid[0][j] == '*' || expandedGrid[0][j] == '.') && !colored[node{0, j}] {
			color(expandedGrid, colored, 0, j)
		}

		if (expandedGrid[m-1][j] == '*' || expandedGrid[m-1][j] == '.') && !colored[node{m - 1, j}] {
			color(expandedGrid, colored, m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		if (expandedGrid[i][0] == '*' || expandedGrid[i][0] == '.') && !colored[node{i, 0}] {
			color(expandedGrid, colored, i, 0)
		}
		if (expandedGrid[i][n-1] == '*' || expandedGrid[i][n-1] == '.') && !colored[node{i, n - 1}] {
			color(expandedGrid, colored, i, n-1)
		}
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if expandedGrid[i][j] == '*' && !colored[node{i, j}] {
				res += 1
			}
		}
	}

	return res
}

func findStart(grid []string) node {
	var start node
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				start.row = i
				start.col = j
			}
		}
	}

	return start
}

func expandGrid(grid []string) []string {
	m := len(grid)
	n := len(grid[0])
	var tmpGrid []string
	for i := 0; i < m; i++ {
		var row strings.Builder
		for j := 0; j < n; j++ {
			row.WriteByte(grid[i][j])
			if j < n-1 {
				row.WriteByte(getHorizontalConnector(grid[i][j], grid[i][j+1]))
			}
		}
		tmpGrid = append(tmpGrid, row.String())
	}

	var expandedGrid []string
	m = len(tmpGrid)
	n = len(tmpGrid[0])

	for i := 0; i < m; i++ {
		var row strings.Builder
		for j := 0; j < n && i < m-1; j++ {
			row.WriteByte(getVerticalConnector(tmpGrid[i][j], tmpGrid[i+1][j]))
		}
		expandedGrid = append(expandedGrid, tmpGrid[i])
		if i < m-1 {
			expandedGrid = append(expandedGrid, row.String())
		}
	}

	return expandedGrid
}

func color(grid []string, colored map[node]bool, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) ||
		colored[node{r, c}] || grid[r][c] != '*' && grid[r][c] != '.' {

		return
	}

	colored[node{r, c}] = true
	for dr := -1; dr < 2; dr++ {
		for dc := -1; dc < 2; dc++ {
			if dr == dc || dr != 0 && dc != 0 {
				continue
			}
			nr, nc := r+dr, c+dc
			color(grid, colored, nr, nc)
		}
	}
}

func buildGraph(start node, grid []string) map[node][]node {
	adj := map[node][]node{}
	m := len(grid)
	n := len(grid[0])
	queue := []node{start}
	visited := map[node]bool{start: true}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for dr := -1; dr < 2; dr++ {
			for dc := -1; dc < 2; dc++ {
				if dr == dc || dr != 0 && dc != 0 {
					continue
				}
				nr, nc := curr.row+dr, curr.col+dc
				next := node{nr, nc}
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}

				if isConnected(grid, curr.row, curr.col, nr, nc) {
					adj[curr] = append(adj[curr], next)
					if !visited[next] {
						queue = append(queue, next)
						visited[next] = true
					}
				}
			}
		}
	}

	return adj
}

func getVerticalConnector(top, bottom byte) byte {
	if top == 'S' {
		if mapping[north][bottom] {
			return '|'
		}
		return '.'
	}
	if !mapping[south][top] {
		return '.'
	}
	return '|'
}

func getHorizontalConnector(left, right byte) byte {
	if left == 'S' {
		if mapping[west][right] {
			return '-'
		}
		return '.'
	}

	if !mapping[east][left] {
		return '.'
	}
	return '-'
}

func isConnected(grid []string, row, col, nextRow, nextCol int) bool {
	if grid[nextRow][nextCol] == '.' {
		return false
	}

	dir := directionLookup[delta{nextRow - row, nextCol - col}]
	nextPipePointsToCurr := grid[nextRow][nextCol] == 'S' || mapping[reverseLookup[dir]][grid[nextRow][nextCol]]
	currPipePointsToNext := grid[row][col] == 'S' || mapping[dir][grid[row][col]]

	return nextPipePointsToCurr && currPipePointsToNext
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
