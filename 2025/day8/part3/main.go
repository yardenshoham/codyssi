package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"unicode"
)

func reduce(line []byte) []byte {
	if len(line) < 2 {
		return line
	}
	for i := range len(line) - 1 {
		if unicode.IsDigit(rune(line[i])) || unicode.IsDigit(rune(line[i+1])) {
			if unicode.IsLetter(rune(line[i])) || unicode.IsLetter(rune(line[i+1])) {
				return slices.Delete(line, i, i+2)
			}
		}
	}
	return line
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	lines := bytes.Fields(input)
	for i := range lines {
		for {
			newLine := reduce(lines[i])
			if len(newLine) == len(lines[i]) {
				break
			}
			lines[i] = newLine
		}
		sum += len(lines[i])
	}
	fmt.Println(sum)
}
