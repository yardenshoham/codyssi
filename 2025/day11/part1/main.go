package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func digit(d byte) int {
	if '0' <= d && d <= '9' {
		return int(d - '0')
	}
	if 'A' <= d && d <= 'Z' {
		return int(d-'A') + 10
	}
	if 'a' <= d && d <= 'z' {
		return int(d-'a') + 36
	}
	panic("bad byte")
}

func number(digits []byte, base int) int {
	result := 0
	slices.Reverse(digits)
	for i, d := range digits {
		result += int(float64(digit(d)) * math.Pow(float64(base), float64(i)))
	}
	return result
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	maxInt := 0
	for _, line := range lines {
		parts := bytes.Fields(line)
		if len(parts) != 2 {
			panic("bad input: " + string(line))
		}
		base, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			panic(err)
		}
		maxInt = max(maxInt, number(parts[0], base))
	}
	fmt.Println(maxInt)
}
