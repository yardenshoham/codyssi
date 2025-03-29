package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/albertorestifo/dijkstra"
)

func cycleCost(graph dijkstra.Graph, visited map[string]struct{}, source, parent, current string) (int, bool) {
	if source == current && parent != "" {
		return 0, true
	}
	if _, ok := visited[current]; ok {
		return 0, false
	}
	visited[current] = struct{}{}
	maxCycle := 0
	for adj, cost := range graph[current] {
		result, valid := cycleCost(graph, visited, source, current, adj)
		if !valid {
			continue
		}
		maxCycle = max(maxCycle, result+cost)
	}
	return maxCycle, maxCycle != 0
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	graph := dijkstra.Graph{}
	for _, line := range strings.Split(string(input), "\n") {
		var from, to string
		var distance int
		_, err := fmt.Sscanf(line, "%s -> %s | %d", &from, &to, &distance)
		if err != nil {
			panic(err)
		}
		if graph[from] == nil {
			graph[from] = map[string]int{}
		}
		if graph[to] == nil {
			graph[to] = map[string]int{}
		}
		graph[from][to] = distance
	}
	maxCycle := 0
	for v := range graph {
		result, valid := cycleCost(graph, map[string]struct{}{}, v, "", v)
		if !valid {
			continue
		}
		maxCycle = max(maxCycle, result)
	}
	fmt.Println(maxCycle)
}
