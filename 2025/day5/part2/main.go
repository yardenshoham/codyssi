package main

import (
	"bytes"
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
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	var zeroPoint image.Point
	minManhattanDistance := math.MaxInt
	var closest image.Point
	for _, line := range lines {
		var p image.Point
		_, err := fmt.Sscanf(string(line), "(%d, %d)", &p.X, &p.Y)
		if err != nil {
			panic(err)
		}
		distance := manhattanDistance(zeroPoint, p)
		if distance > minManhattanDistance {
			continue
		}
		if distance == minManhattanDistance {
			if closest.X < p.X {
				continue
			}
			if closest.X == p.X {
				if closest.Y < p.Y {
					continue
				}
				closest = p
				continue
			}
			closest = p
			continue
		}
		closest = p
		minManhattanDistance = distance
	}
	minManhattanDistance = math.MaxInt
	for _, line := range lines {
		var p image.Point
		_, err := fmt.Sscanf(string(line), "(%d, %d)", &p.X, &p.Y)
		if err != nil {
			panic(err)
		}
		if p == closest {
			continue
		}
		minManhattanDistance = min(minManhattanDistance, manhattanDistance(closest, p))
	}
	fmt.Println(minManhattanDistance)
}
