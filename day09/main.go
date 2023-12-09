package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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
		fmt.Println("part 2:", part2(input))
	}
}

func part1(input string) int {
	dataset := parse(input)
	res := 0

	for i := range dataset {
		res += nextValue(dataset[i])
	}

	return res
}

func part2(input string) int {
	dataset := parse(input)
	for i := range dataset {
		slices.Reverse(dataset[i])
	}
	res := 0

	for i := range dataset {
		res += nextValue(dataset[i])
	}

	return res
}

func nextValue(curr []int) int {
	res := curr[len(curr)-1]
	for !allZeroes(curr) {
		curr = getDeltas(curr)
		res += curr[len(curr)-1]
	}

	return res
}

func allZeroes(nums []int) bool {
	for i := range nums {
		if nums[i] != 0 {
			return false
		}
	}
	return true
}

func getDeltas(nums []int) []int {
	var res []int
	for i := 1; i < len(nums); i++ {
		res = append(res, nums[i]-nums[i-1])
	}
	return res
}

func parse(input string) [][]int {
	var dataset [][]int

	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		var nums []int
		for _, numStr := range strings.Fields(line) {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		dataset = append(dataset, nums)
	}

	return dataset
}
