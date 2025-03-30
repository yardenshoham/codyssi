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

func combination(memo [][]item, items []item, index int, units int) item {
	if index == 0 || units == 0 {
		return item{}
	}
	if memo[index][units].quality != -1 {
		return memo[index][units]
	}
	pick := item{}
	if items[index-1].cost <= units {
		result := combination(memo, items, index-1, units-items[index-1].cost)
		pick.quality = items[index-1].quality + result.quality
		pick.uniqueMaterials = items[index-1].uniqueMaterials + result.uniqueMaterials
	}
	doNotPick := combination(memo, items, index-1, units)
	var result item
	if pick.quality > doNotPick.quality {
		result = pick
	} else if pick.quality < doNotPick.quality {
		result = doNotPick
	} else if pick.uniqueMaterials < doNotPick.uniqueMaterials {
		result = pick
	} else {
		result = doNotPick
	}
	memo[index][units] = result
	return result
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
	units := 300
	memo := make([][]item, len(items)+1)
	for i := range memo {
		memo[i] = make([]item, units+1)
		for j := range memo[i] {
			memo[i][j] = item{quality: -1}
		}
	}
	bestCombination := combination(memo, items, len(items), units)
	fmt.Println(bestCombination.quality * bestCombination.uniqueMaterials)
}
