package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var from1, to1, from2, to2 int
		_, err := fmt.Sscanf(scanner.Text(), "%d-%d %d-%d", &from1, &to1, &from2, &to2)
		if err != nil {
			panic(err)
		}
		if from1 > from2 {
			from1, to1, from2, to2 = from2, to2, from1, to1
		}
		if to1 < from2 {
			sum += to1 - from1 + to2 - from2 + 2
			continue
		}
		sum += max(from1, to1, from2, to2) - min(from1, to1, from2, to2) + 1
	}
	fmt.Println(sum)
}
