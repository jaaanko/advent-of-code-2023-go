package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const size = 256

type node struct {
	key      string
	val      int
	prev     *node
	next     *node
	sentinel bool
}

type customMap struct {
	buckets []node
}

func newMap(n int) *customMap {
	customMap := &customMap{}
	for i := 0; i < n; i++ {
		customMap.buckets = append(customMap.buckets, node{sentinel: true})
	}

	return customMap
}

func (m *customMap) insert(key string, value int) {
	i := hash(key)
	curr := &m.buckets[i]

	for curr != nil {
		if !curr.sentinel && curr.key == key {
			curr.val = value
			return
		}
		if curr.next == nil {
			curr.next = &node{key: key, val: value, prev: curr}
			return
		}
		curr = curr.next
	}
}

func (m *customMap) remove(key string) {
	i := hash(key)
	curr := &m.buckets[i]

	for curr != nil {
		if !curr.sentinel && curr.key == key {
			prev := curr.prev
			next := curr.next
			prev.next = next

			if next != nil {
				next.prev = prev
			}
			curr.prev = nil
			curr.next = nil

			return
		}
		curr = curr.next
	}
}

func (m *customMap) power() int {
	res := 0

	for i := range m.buckets {
		box := i + 1
		slot := 1
		curr := m.buckets[i].next

		for curr != nil {
			res += box * slot * curr.val
			slot += 1
			curr = curr.next
		}

	}

	return res
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
	sequence := parse(input)
	res := 0
	for _, step := range sequence {
		res += hash(step)
	}

	return res
}

func part2(input string) int {
	sequence := parse(input)
	customMap := newMap(size)

	for _, step := range sequence {
		if strings.Contains(step, "=") {
			pair := strings.Split(step, "=")
			val, _ := strconv.Atoi(pair[1])
			customMap.insert(pair[0], val)
		} else {
			customMap.remove(step[0 : len(step)-1])
		}
	}

	return customMap.power()
}

func hash(step string) int {
	res := 0
	for _, char := range step {
		res += int(char)
		res *= 17
		res %= size
	}

	return res
}

func parse(input string) []string {
	return strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), ",")
}
