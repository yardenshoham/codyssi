package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/albertorestifo/dijkstra"
)

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
	costs := []int{}
	for v := range graph {
		_, cost, err := graph.Path("STT", v)
		if err != nil {
			panic(err)
		}
		costs = append(costs, cost)
	}
	slices.Sort(costs)
	largest := len(costs) - 1
	fmt.Println(costs[largest] * costs[largest-1] * costs[largest-2])
}
