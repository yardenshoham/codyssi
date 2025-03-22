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
	values := make([]int, len(input))
	sum := 0
	for i, r := range string(input) {
		if 'a' <= r && r <= 'z' {
			values[i] = int(r - 'a' + 1)
		} else if 'A' <= r && r <= 'Z' {
			values[i] = int(r - 'A' + 27)
		} else {
			newValue := values[i-1]*2 - 5
			for newValue < 1 {
				newValue += 52
			}
			for newValue > 52 {
				newValue -= 52
			}
			values[i] = newValue
		}
		sum += values[i]
	}
	fmt.Println(sum)
}
