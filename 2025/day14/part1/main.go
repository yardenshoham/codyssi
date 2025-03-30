package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type item struct {
	code            string
	quality         int
	cost            int
	uniqueMaterials int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	items := []item{}
	for scanner.Scan() {
		var i item
		var index int
		_, err := fmt.Sscanf(scanner.Text(), "%d %s | Quality : %d, Cost : %d, Unique Materials : %d", &index, &i.code, &i.quality, &i.cost, &i.uniqueMaterials)
		if err != nil {
			panic(err)
		}
		items = append(items, i)
	}
	slices.SortFunc(items, func(a, b item) int {
		if a.quality != b.quality {
			return b.quality - a.quality
		}
		return b.cost - a.cost
	})
	sum := 0
	for _, i := range items[:5] {
		sum += i.uniqueMaterials
	}
	fmt.Println(sum)
}
