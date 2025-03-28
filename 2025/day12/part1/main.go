package main

import (
	"bytes"
	"fmt"
	"image"
	"math"
	"os"
	"strconv"
	"strings"
)

func buildGrid(part0 []byte) map[image.Point]int {
	res := make(map[image.Point]int)
	for y, line := range bytes.Split(part0, []byte{'\n'}) {
		for x, numAsBytes := range bytes.Fields(line) {
			n, err := strconv.Atoi(string(numAsBytes))
			if err != nil {
				panic(err)
			}
			res[image.Pt(x, y)] = n
		}
	}
	return res
}

func gridSize(grid map[image.Point]int) int {
	return int(math.Sqrt(float64(len(grid))))
}

func shift(grid map[image.Point]int, amount int, row bool, index int) {
	n := gridSize(grid)
	if row {
		newRow := make([]int, n)
		for x := range n {
			newRow[(x+amount)%n] = grid[image.Pt(x, index)]
		}
		for i := range n {
			grid[image.Pt(i, index)] = newRow[i]
		}
		return
	}
	newCol := make([]int, n)
	for y := range n {
		newCol[(y+amount)%n] = grid[image.Pt(index, y)]
	}
	for i := range n {
		grid[image.Pt(index, i)] = newCol[i]
	}
}

func fix(n int) int {
	for n < 0 {
		n += 1073741824
	}
	for n > 1073741823 {
		n -= 1073741824
	}
	return n
}

func instruction(grid map[image.Point]int, f func(int, int) int, amount int, row bool, index int) {
	n := gridSize(grid)
	if row {
		for x := range n {
			grid[image.Pt(x, index)] = fix(f(grid[image.Pt(x, index)], amount))
		}
		return
	}
	for y := range n {
		grid[image.Pt(index, y)] = fix(f(grid[image.Pt(index, y)], amount))
	}
}

func instructionAll(grid map[image.Point]int, f func(int, int) int, amount int) {
	n := gridSize(grid)
	for i := range n {
		instruction(grid, f, amount, false, i)
	}
}

func executeInstructions(grid map[image.Point]int, instructions []byte) {
	for _, inst := range strings.Split(string(instructions), "\n") {
		parts := strings.Fields(inst)
		if len(parts) == 5 {
			row := parts[1] == "ROW"
			index, err := strconv.Atoi(parts[2])
			if err != nil {
				panic(err)
			}
			index--
			amount, err := strconv.Atoi(parts[4])
			if err != nil {
				panic(err)
			}
			shift(grid, amount, row, index)
			continue
		}
		var op func(int, int) int
		switch parts[0] {
		case "ADD":
			op = func(i1, i2 int) int { return i1 + i2 }
		case "SUB":
			op = func(i1, i2 int) int { return i1 - i2 }
		case "MULTIPLY":
			op = func(i1, i2 int) int { return i1 * i2 }
		default:
			panic("bad op: " + parts[0])
		}
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		if len(parts) == 4 {
			row := parts[2] == "ROW"
			index, err := strconv.Atoi(parts[3])
			if err != nil {
				panic(err)
			}
			index--
			instruction(grid, op, amount, row, index)
			continue
		}
		instructionAll(grid, op, amount)
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := bytes.Split(input, []byte{'\n', '\n'})
	if len(parts) != 3 {
		panic("should have 3 parts")
	}
	grid := buildGrid(parts[0])
	executeInstructions(grid, parts[1])
	maxSum := 0
	n := gridSize(grid)
	for y := range n {
		sumRow := 0
		sumCol := 0
		for x := range n {
			sumRow += grid[image.Pt(x, y)]
			sumCol += grid[image.Pt(y, x)]
		}
		maxSum = max(maxSum, sumRow, sumCol)
	}
	fmt.Println(maxSum)
}
