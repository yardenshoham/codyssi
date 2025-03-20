package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		line := scanner.Bytes()
		keep := len(line) / 10
		for i := range keep {
			sum += int(line[i] - 'A' + 1)
		}
		for i := len(line) - 1; i > len(line)-1-keep; i-- {
			sum += int(line[i] - 'A' + 1)
		}
		left := len(line) - 2*keep
		for _, r := range strconv.Itoa(left) {
			sum += int(r - '0')
		}
	}
	fmt.Println(sum)
}
