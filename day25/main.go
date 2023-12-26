package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type edge struct {
	u, v string
}

func main() {
	p1 := flag.Bool("p1", false, "run part 1")
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
}

func part1(input string) int {
	adj := parse(input)
	seenEdges := map[edge]bool{}
	edges := []edge{}

	for u := range adj {
		for nei := range adj[u] {
			if seenEdges[edge{u, nei}] || seenEdges[edge{nei, u}] {
				continue
			}
			seenEdges[edge{u, nei}] = true
			edges = append(edges, edge{u, nei})
		}
	}

	var edgesToCut [3]edge
	var root string
	for k := range adj {
		root = k
		break
	}

	for i := 0; i < len(edges); i++ {
		found := false
		for j := i + 1; j < len(edges); j++ {
			removeEdge(edges[i], adj)
			removeEdge(edges[j], adj)
			bridge := findBridge(root, "", 1, map[string]int{}, map[string]int{}, adj, map[string]bool{})
			addEdge(edges[i], adj)
			addEdge(edges[j], adj)

			if bridge != nil {
				found = true
				edgesToCut = [3]edge{edges[i], edges[j], *bridge}
			}
		}
		if found {
			break
		}
	}

	for _, edge := range edgesToCut {
		removeEdge(edge, adj)
	}

	res := 1
	visited := map[string]bool{}
	for k := range adj {
		if visited[k] {
			continue
		}
		res *= count(k, adj, visited)
	}

	return res
}

func removeEdge(toRemove edge, adj map[string]map[string]bool) {
	delete(adj[toRemove.u], toRemove.v)
	delete(adj[toRemove.v], toRemove.u)
}

func addEdge(toAdd edge, adj map[string]map[string]bool) {
	adj[toAdd.u][toAdd.v] = true
	adj[toAdd.v][toAdd.u] = true
}

func findBridge(u, p string, t int, entryTime, low map[string]int, adj map[string]map[string]bool, visited map[string]bool) *edge {
	visited[u] = true
	entryTime[u] = t
	low[u] = t

	for nei := range adj[u] {
		if nei == p {
			continue
		}
		if visited[nei] {
			low[u] = min(low[u], entryTime[nei])
		} else {
			bridge := findBridge(nei, u, t+1, entryTime, low, adj, visited)
			if bridge != nil {
				return bridge
			}
			low[u] = min(low[u], low[nei])
			if low[nei] > entryTime[u] {
				return &edge{u, nei}
			}
		}
	}

	return nil
}

func count(u string, adj map[string]map[string]bool, visited map[string]bool) int {
	visited[u] = true
	res := 1

	for nei := range adj[u] {
		if visited[nei] {
			continue
		}
		res += count(nei, adj, visited)
	}

	return res
}

func parse(input string) map[string]map[string]bool {
	adj := map[string]map[string]bool{}
	for _, line := range strings.Split(strings.ReplaceAll(input, "\r\n", "\n"), "\n") {
		line := strings.Split(line, ": ")
		u := line[0]
		if adj[u] == nil {
			adj[u] = map[string]bool{}
		}

		for _, nei := range strings.Fields(line[1]) {
			adj[u][nei] = true
			if adj[nei] == nil {
				adj[nei] = map[string]bool{}
			}
			adj[nei][u] = true
		}
	}
	return adj
}
