package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
)

func manhattanDistance(a, b image.Point) int {
	diff := a.Sub(b)
	if diff.X < 0 {
		diff.X *= -1
	}
	if diff.Y < 0 {
		diff.Y *= -1
	}
	return diff.X + diff.Y
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var zeroPoint image.Point
	minManhattanDistance := math.MaxInt
	maxManhattanDistance := 0
	for scanner.Scan() {
		var p image.Point
		_, err := fmt.Sscanf(scanner.Text(), "(%d, %d)", &p.X, &p.Y)
		if err != nil {
			panic(err)
		}
		distance := manhattanDistance(zeroPoint, p)
		maxManhattanDistance = max(maxManhattanDistance, distance)
		minManhattanDistance = min(minManhattanDistance, distance)
	}
	fmt.Println(maxManhattanDistance - minManhattanDistance)
}
