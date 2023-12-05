package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type interval struct {
	left, right int
}

type mapEntry struct {
	dest, source interval
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
	starts, mappings := parse(input, true)
	res := math.MaxInt64

	for _, interval := range starts {
		layers := createSortedLayers(mappings)
		min := calculateMin(layers, mappings, interval)
		if min < res {
			res = min
		}
	}

	return res
}

func part2(input string) int {
	starts, mappings := parse(input, false)
	res := math.MaxInt64

	for _, interval := range starts {
		layers := createSortedLayers(mappings)
		min := calculateMin(layers, mappings, interval)
		if min < res {
			res = min
		}
	}

	return res
}

func createSortedLayers(mappings [][]mapEntry) [][]interval {
	var layers [][]interval
	for _, mapping := range mappings {
		var layer []interval
		for _, entry := range mapping {
			layer = append(layer, entry.source)
		}
		sort.Slice(layer, func(i, j int) bool {
			return layer[i].left < layer[j].left
		})
		layers = append(layers, layer)
	}

	return layers
}

func calculateMin(layers [][]interval, mappings [][]mapEntry, initial interval) int {
	currIntervals := []interval{initial}
	for i, mapping := range mappings {
		var newIntervals []interval
		layer := layers[i]

		for _, upper := range currIntervals {
			curr := upper
			appendEnd := true

			for _, lower := range layer {
				if curr.right < lower.left || lower.right < curr.left {
					continue
				}

				left, middle, right := split(curr, lower)
				if left != nil {
					newIntervals = append(newIntervals, *left)
				}
				if middle != nil {
					newIntervals = append(newIntervals, *middle)
				}
				if right == nil {
					appendEnd = false
					break
				}
				curr = *right
			}

			if appendEnd {
				newIntervals = append(newIntervals, curr)
			}
		}

		if len(newIntervals) > 0 {
			currIntervals = newIntervals
			for i := range currIntervals {
				for _, entry := range mapping {
					if entry.source.left <= currIntervals[i].left && currIntervals[i].right <= entry.source.right {
						currIntervals[i] = translate(entry, currIntervals[i])
						break
					}
				}
			}
		}
	}

	res := math.MaxInt64
	for _, interval := range currIntervals {
		if interval.left < res {
			res = interval.left
		}
	}

	return res
}

func translate(entry mapEntry, in interval) interval {
	newLeft := in.left + (entry.dest.left - entry.source.left)
	newRight := in.right + (entry.dest.right - entry.source.right)
	return interval{newLeft, newRight}
}

func split(toBeSplit interval, in interval) (left, middle, right *interval) {
	if toBeSplit.left >= in.left && toBeSplit.right <= in.right {
		return nil, &interval{toBeSplit.left, toBeSplit.right}, nil
	}
	if toBeSplit.left < in.left {
		left = &interval{left: toBeSplit.left, right: in.left - 1}
		if toBeSplit.right <= in.right {
			middle = &interval{in.left, toBeSplit.right}
		} else {
			middle = &interval{in.left, in.right}
			right = &interval{in.right + 1, toBeSplit.right}
		}
		return left, middle, right
	}

	middle = &interval{toBeSplit.left, in.right}
	right = &interval{in.right + 1, toBeSplit.right}
	return left, middle, right
}

func parse(input string, p1 bool) ([]interval, [][]mapEntry) {
	var starts []interval
	var mappings [][]mapEntry

	for i, block := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n\n") {
		if i == 0 {
			f := func(r rune) bool { return r == ':' || r == ' ' }
			fields := strings.FieldsFunc(block, f)
			if p1 {
				for j := 1; j < len(fields); j++ {
					left, _ := strconv.Atoi(fields[j])
					starts = append(starts, interval{left: left, right: left})
				}
			} else {
				for j := 1; j < len(fields); j += 2 {
					left, _ := strconv.Atoi(fields[j])
					len, _ := strconv.Atoi(fields[j+1])
					starts = append(starts, interval{left: left, right: left + len - 1})
				}
			}
		} else {
			var mapping []mapEntry
			lines := strings.Split(block, "\n")
			for j := 1; j < len(lines); j++ {
				var dest, source, rangeLen int
				fmt.Sscanf(lines[j], "%d %d %d", &dest, &source, &rangeLen)
				mapping = append(mapping, mapEntry{
					dest:   interval{left: dest, right: dest + rangeLen - 1},
					source: interval{left: source, right: source + rangeLen - 1},
				})
			}
			mappings = append(mappings, mapping)
		}
	}

	return starts, mappings
}
