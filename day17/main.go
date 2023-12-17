package main

import (
	"container/heap"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type direction struct {
	row, col int
}

var (
	left  = direction{row: 0, col: -1}
	right = direction{row: 0, col: 1}
	up    = direction{row: -1, col: 0}
	down  = direction{row: 1, col: 0}
)

var rotations = map[direction][]direction{
	left:  {up, down},
	right: {up, down},
	up:    {left, right},
	down:  {left, right},
}

var reverse = map[direction]direction{
	left:  right,
	right: left,
	up:    down,
	down:  up,
}

type state struct {
	row   int
	col   int
	dir   direction
	moves int
}

type item struct {
	heatLoss int
	state    state
	index    int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
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
	return dijkstra(grid, 3, 0)
}

func part2(input string) int {
	grid := parse(input)
	return dijkstra(grid, 10, 4)
}

func dijkstra(grid []string, maxConsecutive, movesNeededBeforeTurn int) int {
	m := len(grid)
	n := len(grid[0])
	startRight := state{row: 0, col: 0, dir: right, moves: 0}
	startDown := state{row: 0, col: 0, dir: down, moves: 0}
	pq := priorityQueue{
		&item{heatLoss: 0, state: startRight, index: 0},
		&item{heatLoss: 0, state: startDown, index: 1},
	}

	minCost := map[state]int{startRight: 0, startDown: 0}
	heap.Init(&pq)

	for len(pq) > 0 {
		curr := heap.Pop(&pq).(*item)
		if minCost[curr.state] < curr.heatLoss {
			continue
		}

		if curr.state.row == m-1 && curr.state.col == n-1 && curr.state.moves >= movesNeededBeforeTurn {
			return curr.heatLoss
		}

		for _, dir := range [4]direction{left, right, up, down} {
			if curr.state.moves == maxConsecutive && !slices.Contains(rotations[curr.state.dir], dir) || dir == reverse[curr.state.dir] {
				continue
			}

			ni, nj := curr.state.row+dir.row, curr.state.col+dir.col
			nextMoves := curr.state.moves

			if curr.state.moves < movesNeededBeforeTurn {
				if dir != curr.state.dir {
					continue
				}
				nextMoves += 1
			} else {
				if dir != curr.state.dir {
					nextMoves = 1
				} else {
					nextMoves = nextMoves%maxConsecutive + 1
				}
			}

			if ni < 0 || ni >= m || nj < 0 || nj >= n {
				continue
			}

			nextState := state{row: ni, col: nj, moves: nextMoves, dir: dir}
			nextHeatLoss := int(rune(grid[ni][nj]) - '0')
			if _, ok := minCost[nextState]; ok && minCost[nextState] <= curr.heatLoss+nextHeatLoss {
				continue
			}

			minCost[nextState] = curr.heatLoss + nextHeatLoss
			heap.Push(&pq, &item{heatLoss: curr.heatLoss + nextHeatLoss, state: nextState})
		}
	}

	return -1
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
