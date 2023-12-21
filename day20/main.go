package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type pulse int

const (
	low pulse = iota
	high
	none
)

type module interface {
	convert(string, pulse) pulse
}

type conjunction struct {
	memory map[string]pulse
}

func (c *conjunction) convert(key string, p pulse) pulse {
	c.memory[key] = p
	for k := range c.memory {
		if c.memory[k] == low {
			return high
		}
	}

	return low
}

type flipFlop struct {
	on bool
}

func (f *flipFlop) convert(_ string, p pulse) pulse {
	if p == high {
		return none
	}

	if !f.on {
		f.on = true
		return high
	}

	f.on = false
	return low
}

type broadcaster struct{}

func (b *broadcaster) convert(_ string, p pulse) pulse {
	return p
}

type state struct {
	moduleName string
	pulse      pulse
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
	adj, moduleLookup := parse(input)
	presses := 1
	totalLow := 0
	totalHigh := 0

	for presses <= 1000 {
		low, high, _ := bfs(adj, moduleLookup, "")
		totalLow += low
		totalHigh += high
		presses += 1
	}

	return totalLow * totalHigh
}

func part2(input string) int {
	res := 1
	// Based on my input, these modules have to emit a low pulse in order
	// for the pulse that eventually reaches "rx" to also be low.
	for _, target := range [4]string{"qx", "db", "gf", "vc"} {
		adj, moduleLookup := parse(input)
		presses := 1
		for {
			_, _, targetEmitsLowPulse := bfs(adj, moduleLookup, target)
			if targetEmitsLowPulse {
				res *= presses
				break
			}
			presses += 1
		}
	}

	return res
}

func bfs(adj map[string][]string, moduleLookup map[string]module, target string) (int, int, bool) {
	start := "broadcaster"
	queue := []state{{moduleName: start, pulse: moduleLookup[start].convert("button", low)}}
	lowCount := 1
	highCount := 0
	targetEmitsLowPulse := false

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.moduleName == target && curr.pulse == low {
			targetEmitsLowPulse = true
		}

		for _, nei := range adj[curr.moduleName] {
			if curr.pulse == low {
				lowCount += 1
			} else {
				highCount += 1
			}

			nextModule, ok := moduleLookup[nei]
			if !ok {
				continue
			}

			nextPulse := nextModule.convert(curr.moduleName, curr.pulse)
			if nextPulse != none {
				queue = append(queue, state{moduleName: nei, pulse: nextPulse})
			}
		}
	}

	return lowCount, highCount, targetEmitsLowPulse
}

func parse(input string) (map[string][]string, map[string]module) {
	adj := map[string][]string{}
	moduleLookup := map[string]module{}

	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		line := strings.Split(line, " -> ")
		var module module
		var name string

		if line[0][0] == '%' {
			module = &flipFlop{}
			name = line[0][1:]
		} else if line[0][0] == '&' {
			module = &conjunction{memory: map[string]pulse{}}
			name = line[0][1:]
		} else {
			module = &broadcaster{}
			name = line[0]
		}

		adj[name] = strings.Split(line[1], ", ")
		moduleLookup[name] = module
	}

	for k := range adj {
		for _, nei := range adj[k] {
			if _, ok := moduleLookup[nei].(*conjunction); ok {
				moduleLookup[nei].(*conjunction).memory[k] = low
			}
		}
	}

	return adj, moduleLookup
}
