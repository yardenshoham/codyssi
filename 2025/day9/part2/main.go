package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := bytes.Split(input, []byte{'\n', '\n'})
	if len(parts) != 2 {
		panic("bad input")
	}
	balances := make(map[string]int)
	for _, initialBalance := range bytes.Split(parts[0], []byte{'\n'}) {
		var balance int
		var official string
		_, err := fmt.Sscanf(string(initialBalance), "%s HAS %d", &official, &balance)
		if err != nil {
			panic(err)
		}
		balances[official] = balance
	}
	for _, transaction := range bytes.Split(parts[1], []byte{'\n'}) {
		var from, to string
		var amount int
		_, err := fmt.Sscanf(string(transaction), "FROM %s TO %s AMT %d", &from, &to, &amount)
		if err != nil {
			panic(err)
		}
		if balances[from] < amount {
			amount = balances[from]
		}
		balances[from] -= amount
		balances[to] += amount
	}
	balanceValues := make([]int, 0, len(balances))
	for _, bv := range balances {
		balanceValues = append(balanceValues, bv)
	}
	slices.Sort(balanceValues)
	fmt.Println(balanceValues[len(balanceValues)-1] + balanceValues[len(balanceValues)-2] + balanceValues[len(balanceValues)-3])
}
