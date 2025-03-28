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

func base68digit(n int) byte {
	if 0 <= n && n <= 9 {
		return byte(n + '0')
	}
	if 10 <= n && n <= 35 {
		return byte(n + 'A' - 10)
	}
	if 36 <= n && n <= 61 {
		return byte(n + 'a' - 36)
	}
	switch n {
	case 62:
		return '!'
	case 63:
		return '@'
	case 64:
		return '#'
	case 65:
		return '$'
	case 66:
		return '%'
	case 67:
		return '^'
	}
	panic("bad base68 digit")
}

func base68number(n int) []byte {
	res := []byte{}
	for n != 0 {
		res = append(res, base68digit(n%68))
		n /= 68
	}
	slices.Reverse(res)
	return res
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	sum := 0
	for _, line := range lines {
		parts := bytes.Fields(line)
		if len(parts) != 2 {
			panic("bad input: " + string(line))
		}
		base, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			panic(err)
		}
		sum += number(parts[0], base)
	}
	fmt.Println(string(base68number(sum)))
}
