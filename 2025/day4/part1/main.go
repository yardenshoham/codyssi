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
		for _, b := range scanner.Bytes() {
			sum += int(b - 'A' + 1)
		}
	}
	fmt.Println(sum)
}
