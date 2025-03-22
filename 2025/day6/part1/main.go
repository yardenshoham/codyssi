package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	sum := 0
	for _, r := range string(input) {
		if unicode.IsLetter(r) {
			sum++
		}
	}
	fmt.Println(sum)
}
