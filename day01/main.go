package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
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
	lines := parse(input)
	res := 0

	for _, line := range lines {
		var first, last int
		for _, char := range line {
			if unicode.IsDigit(rune(char)) {
				first = int(rune(char) - '0')
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = int(rune(line[i]) - '0')
				break
			}
		}
		res += first*10 + last
	}

	return res
}

func part2(input string) int {
	lines := parse(input)
	lookup := map[string]int{
		"1": 1, "2": 2, "3": 3,
		"4": 4, "5": 5, "6": 6,
		"7": 7, "8": 8, "9": 9,
		"one": 1, "two": 2, "three": 3,
		"four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9,
	}

	res := 0
	for _, line := range lines {
		var first, last int
		min := len(line)
		for strRep, val := range lookup {
			i := strings.Index(line, strRep)
			if i != -1 && i < min {
				min = i
				first = val
			}
		}

		max := -1
		for strRep, val := range lookup {
			i := strings.LastIndex(line, strRep)
			if i != -1 && i > max {
				max = i
				last = val
			}
		}
		res += first*10 + last
	}

	return res
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n")
}
