package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func lastNumber(line []byte) float64 {
	parts := bytes.Fields(line)
	num, err := strconv.Atoi(string(parts[len(parts)-1]))
	if err != nil {
		panic(err)
	}
	return float64(num)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	add := lastNumber(lines[0])
	multiply := lastNumber(lines[1])
	raise := lastNumber(lines[2])

	numbers := make([]float64, 0, len(lines)-3)
	for i := 4; i < len(lines); i++ {
		numbers = append(numbers, lastNumber(lines[i]))
	}

	slices.Sort(numbers)
	median := numbers[len(numbers)/2]
	fmt.Printf("%.f\n", math.Pow(median, raise)*multiply+add)
}
