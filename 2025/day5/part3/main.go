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

func next(points map[image.Point]struct{}, from image.Point) (int, image.Point) {
	minManhattanDistance := math.MaxInt
	var closest image.Point
	for p := range points {
		distance := manhattanDistance(from, p)
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
	return minManhattanDistance, closest
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	points := make(map[image.Point]struct{})
	for scanner.Scan() {
		var p image.Point
		_, err := fmt.Sscanf(scanner.Text(), "(%d, %d)", &p.X, &p.Y)
		if err != nil {
			panic(err)
		}
		points[p] = struct{}{}
	}
	sum := 0
	var current image.Point
	for len(points) > 0 {
		distance, nextPoint := next(points, current)
		sum += distance
		current = nextPoint
		delete(points, nextPoint)
	}
	fmt.Println(sum)
}
