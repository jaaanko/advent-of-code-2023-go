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

type cube struct {
	x1, y1, z1, x2, y2, z2 int
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
	cubes := parse(input)
	sort.Slice(cubes, func(i, j int) bool {
		return cubes[i].z1 < cubes[j].z1
	})

	supportedBy, numDependencies := createGraph(cubes)
	res := 0

	for i := 0; i < len(cubes); i++ {
		remove := true
		for _, j := range supportedBy[i] {
			if numDependencies[j]-1 == 0 {
				remove = false
				break
			}
		}
		if remove {
			res += 1
		}
	}

	return res
}

func createGraph(cubes []cube) ([][]int, []int) {
	zLookup := map[int]map[int]bool{}
	n := len(cubes)
	numDependencies := make([]int, n)
	supportedBy := make([][]int, n)

	for i, cube := range cubes {
		if zLookup[cube.z2] == nil {
			zLookup[cube.z2] = map[int]bool{}
		}
		zLookup[cube.z2][i] = true
	}

	for i := 0; i < n; i++ {
		var nextCube = cubes[i]
		collisions := findCollisions(nextCube.z1-1, cubes, zLookup, nextCube)
		for nextCube.z1-1 > 0 && len(collisions) == 0 {
			nextCube = cube{nextCube.x1, nextCube.y1, nextCube.z1 - 1, nextCube.x2, nextCube.y2, nextCube.z2 - 1}
			collisions = findCollisions(nextCube.z1-1, cubes, zLookup, nextCube)
		}

		delete(zLookup[cubes[i].z2], i)
		if zLookup[nextCube.z2] == nil {
			zLookup[nextCube.z2] = map[int]bool{}
		}

		zLookup[nextCube.z2][i] = true
		cubes[i] = nextCube
		numDependencies[i] = len(collisions)

		for _, j := range collisions {
			supportedBy[j] = append(supportedBy[j], i)
		}
	}

	return supportedBy, numDependencies
}

func findCollisions(z int, cubes []cube, zLookup map[int]map[int]bool, c cube) []int {
	var collisions []int
	for i := range zLookup[z] {
		if (cubes[i].x1 <= c.x1 && c.x1 <= cubes[i].x2 || c.x1 <= cubes[i].x1 && cubes[i].x1 <= c.x2) &&
			(cubes[i].y1 <= c.y1 && c.y1 <= cubes[i].y2 || c.y1 <= cubes[i].y1 && cubes[i].y1 <= c.y2) {

			collisions = append(collisions, i)
		}
	}

	return collisions
}

func part2(input string) int {
	cubes := parse(input)
	sort.Slice(cubes, func(i, j int) bool {
		return cubes[i].z1 < cubes[j].z1
	})
	supportedBy, numDependencies := createGraph(cubes)
	n := len(cubes)
	res := 0

	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		copy(tmp, numDependencies)
		res += count(i, tmp, supportedBy)
	}

	return res
}

func count(i int, numDependencies []int, supportedBy [][]int) int {
	queue := []int{i}
	res := 0

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		res += 1
		for _, j := range supportedBy[curr] {
			numDependencies[j] -= 1
			if numDependencies[j] == 0 {
				queue = append(queue, j)
			}
		}
	}

	return res - 1
}

func parse(input string) []cube {
	var cubes []cube
	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		line := strings.Split(line, "~")
		coords1 := strings.Split(line[0], ",")
		coords2 := strings.Split(line[1], ",")

		x1, _ := strconv.Atoi(coords1[0])
		y1, _ := strconv.Atoi(coords1[1])
		z1, _ := strconv.Atoi(coords1[2])

		x2, _ := strconv.Atoi(coords2[0])
		y2, _ := strconv.Atoi(coords2[1])
		z2, _ := strconv.Atoi(coords2[2])

		cubes = append(cubes, cube{x1, y1, z1, x2, y2, z2})
	}

	return cubes
}
