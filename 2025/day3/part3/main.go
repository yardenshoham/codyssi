package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(input), "\n")
	best := 0
	for i := 0; i < len(lines)-1; i++ {
		set := make(map[int]struct{})
		var from1, to1, from2, to2 int
		_, err := fmt.Sscanf(lines[i], "%d-%d %d-%d", &from1, &to1, &from2, &to2)
		if err != nil {
			panic(err)
		}
		for j := from1; j <= to1; j++ {
			set[j] = struct{}{}
		}
		for j := from2; j <= to2; j++ {
			set[j] = struct{}{}
		}
		_, err = fmt.Sscanf(lines[i+1], "%d-%d %d-%d", &from1, &to1, &from2, &to2)
		if err != nil {
			panic(err)
		}
		for j := from1; j <= to1; j++ {
			set[j] = struct{}{}
		}
		for j := from2; j <= to2; j++ {
			set[j] = struct{}{}
		}
		best = max(len(set), best)
	}
	fmt.Println(best)
}
