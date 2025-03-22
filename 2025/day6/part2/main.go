package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, r := range string(input) {
		if 'a' <= r && r <= 'z' {
			sum += int(r - 'a' + 1)
		}
		if 'A' <= r && r <= 'Z' {
			sum += int(r - 'A' + 27)
		}
	}
	fmt.Println(sum)
}
