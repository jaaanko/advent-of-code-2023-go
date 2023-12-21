package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type interval struct {
	left, right int
}

type step struct {
	predicate    func(int) bool
	param        rune
	next         string
	operator     string
	rightOperand int
}

type workflow struct {
	steps       []step
	defaultDest string
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
	workflowLookup, parts := parse(input)
	res := 0
	for _, part := range parts {
		if accepted("in", part, workflowLookup) {
			res += part['x'] + part['m'] + part['a'] + part['s']
		}
	}
	return res
}

func accepted(curr string, part map[rune]int, workflowLookup map[string]workflow) bool {
	workflow, ok := workflowLookup[curr]
	if !ok {
		return curr == "A"
	}

	for _, step := range workflow.steps {
		if step.predicate(part[step.param]) {
			return accepted(step.next, part, workflowLookup)
		}
	}

	return accepted(workflow.defaultDest, part, workflowLookup)
}

func part2(input string) int {
	workflowLookup, _ := parse(input)
	constraints := map[rune]interval{
		'x': {left: 1, right: 4000},
		'm': {left: 1, right: 4000},
		'a': {left: 1, right: 4000},
		's': {left: 1, right: 4000},
	}

	return countWays("in", constraints, workflowLookup)
}

func countWays(curr string, constraints map[rune]interval, workflowLookup map[string]workflow) int {
	workflow, ok := workflowLookup[curr]
	if !ok {
		if curr == "A" {
			prod := 1
			for _, char := range "xmas" {
				prod *= constraints[char].right - constraints[char].left + 1
			}
			return prod
		}
		return 0
	}

	copy := map[rune]interval{}
	for k, v := range constraints {
		copy[k] = v
	}

	ways := 0
	defaultCase := true
	for _, step := range workflow.steps {
		prev := copy[step.param]
		alwaysTrue := false

		if step.operator == ">" {
			if prev.left > step.rightOperand {
				alwaysTrue = true
			} else {
				copy[step.param] = interval{left: step.rightOperand + 1, right: prev.right}
				prev.right = step.rightOperand
			}

		} else {
			if prev.right < step.rightOperand {
				alwaysTrue = true
			} else {
				copy[step.param] = interval{left: prev.left, right: step.rightOperand - 1}
				prev.left = step.rightOperand
			}
		}

		ways += countWays(step.next, copy, workflowLookup)
		copy[step.param] = prev
		if alwaysTrue {
			defaultCase = false
			break
		}
	}

	if defaultCase {
		ways += countWays(workflowLookup[curr].defaultDest, copy, workflowLookup)
	}

	return ways
}

func parse(input string) (map[string]workflow, []map[rune]int) {
	blocks := strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n")
	workflowLookup := map[string]workflow{}
	var parts = []map[rune]int{}

	for _, line := range strings.Split(blocks[0], "\n") {
		info := strings.Split(line, "{")
		steps := info[1][0 : len(info[1])-1]
		workflow := workflow{}

		for _, condition := range strings.Split(steps, ",") {
			if !strings.Contains(condition, ":") {
				workflow.defaultDest = condition
			} else {
				step := step{}
				stepInfo := strings.Split(condition, ":")
				step.next = stepInfo[1]
				step.param = rune(stepInfo[0][0])
				m, _ := strconv.Atoi(stepInfo[0][2:])
				step.rightOperand = m
				step.operator = string(stepInfo[0][1])

				if stepInfo[0][1] == '>' {
					step.predicate = func(n int) bool { return n > m }
				} else {
					step.predicate = func(n int) bool { return n < m }
				}
				workflow.steps = append(workflow.steps, step)
			}
		}
		workflowLookup[info[0]] = workflow
	}

	for _, line := range strings.Split(blocks[1], "\n") {
		part := map[rune]int{}
		for _, category := range strings.Split(line[1:len(line)-1], ",") {
			num, _ := strconv.Atoi(category[2:])
			part[rune(category[0])] = num
		}
		parts = append(parts, part)
	}

	return workflowLookup, parts
}
