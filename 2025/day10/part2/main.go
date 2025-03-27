package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

func safest(grid map[image.Point]int, visited map[image.Point]int, start, end image.Point) int {
	if res, ok := visited[start]; ok {
		return res
	}
	dangerLevel, ok := grid[start]
	if !ok {
		return 1000
	}
	if start == end {
		return dangerLevel
	}
	res := dangerLevel + min(safest(grid, visited, start.Add(image.Pt(0, 1)), end), safest(grid, visited, start.Add(image.Pt(1, 0)), end))
	visited[start] = res
	return res
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	grid := make(map[image.Point]int, len(lines)*len(lines))
	for y, line := range lines {
		for x, b := range bytes.Fields(line) {
			grid[image.Pt(x, y)] = int(b[0] - '0')
		}
	}
	fmt.Println(safest(grid, make(map[image.Point]int), image.Pt(0, 0), image.Pt(14, 14)))
}
