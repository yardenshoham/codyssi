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
		line := scanner.Bytes()
		cur := line[0]
		count := 1
		for _, b := range line[1:] {
			if cur != b {
				sum += int(cur - 'A' + 1)
				for count > 0 {
					sum += count % 10
					count /= 10
				}
				cur = b
				count = 1
				continue
			}
			count++
		}
		sum += int(cur - 'A' + 1)
		for count > 0 {
			sum += count % 10
			count /= 10
		}
	}
	fmt.Println(sum)
}
