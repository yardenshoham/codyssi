package main

import (
	"bufio"
	"fmt"
	"os"
)

type item struct {
	code            string
	quality         int
	cost            int
	uniqueMaterials int
}

func combination(items []item, index int, units int) item {
	if index == 0 || units == 0 {
		return item{}
	}
	pick := item{}
	if items[index-1].cost <= units {
		result := combination(items, index-1, units-items[index-1].cost)
		pick.quality = items[index-1].quality + result.quality
		pick.uniqueMaterials = items[index-1].uniqueMaterials + result.uniqueMaterials
	}
	doNotPick := combination(items, index-1, units)
	if pick.quality > doNotPick.quality {
		return pick
	} else if pick.quality < doNotPick.quality {
		return doNotPick
	}
	if pick.uniqueMaterials < doNotPick.uniqueMaterials {
		return pick
	}
	return doNotPick
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
	bestCombination := combination(items, len(items), 30)
	fmt.Println(bestCombination.quality * bestCombination.uniqueMaterials)
}
