package main

import (
	"bytes"
	"fmt"
	"image"
	"math"
	"os"
)

func dangerLevel(grid map[image.Point]int, start image.Point, step image.Point) int {
	current := start
	sum := 0
	for {
		sum += grid[current]
		current = current.Add(step)
		if _, ok := grid[current]; !ok {
			break
		}
	}
	return sum
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
	minimumDangerLevel := math.MaxInt
	for i := range len(lines) {
		minimumDangerLevel = min(minimumDangerLevel, dangerLevel(grid, image.Pt(0, i), image.Pt(1, 0)), dangerLevel(grid, image.Pt(i, 0), image.Pt(0, 1)))
	}
	fmt.Println(minimumDangerLevel)
}
