package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Fields(input)
	sum := int(lines[0][0]-'0')*10 + int(lines[1][0]-'0')
	for i := 2; i < len(lines)-1; i += 2 {
		multiplier := 1
		if lines[len(lines)-1][len(lines)-3-i/2+1] == '-' {
			multiplier = -1
		}
		sum += multiplier * (int(lines[i][0]-'0')*10 + int(lines[i+1][0]-'0'))
	}
	fmt.Println(sum)
}
