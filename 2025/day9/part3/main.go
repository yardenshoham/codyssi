package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
)

type debt struct {
	owedTo string
	amount int
}

func executeTransaction(balances map[string]int, debts map[string][]debt, from, to string, amount int) {
	if balances[from] < amount {
		debts[from] = append(debts[from], debt{to, amount - balances[from]})
		amount = balances[from]
	}
	balances[from] -= amount
	balances[to] += amount
	for i, d := range debts[to] {
		if balances[to] == 0 {
			break
		}
		if d.amount == 0 {
			continue
		}
		toPay := min(balances[to], d.amount)
		debts[to][i].amount = d.amount - toPay
		executeTransaction(balances, debts, to, d.owedTo, toPay)
	}
}

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
	debts := make(map[string][]debt)
	for _, transaction := range bytes.Split(parts[1], []byte{'\n'}) {
		var from, to string
		var amount int
		_, err := fmt.Sscanf(string(transaction), "FROM %s TO %s AMT %d", &from, &to, &amount)
		if err != nil {
			panic(err)
		}
		executeTransaction(balances, debts, from, to, amount)
	}
	balanceValues := make([]int, 0, len(balances))
	for _, bv := range balances {
		balanceValues = append(balanceValues, bv)
	}
	slices.Sort(balanceValues)
	fmt.Println(balanceValues[len(balanceValues)-1] + balanceValues[len(balanceValues)-2] + balanceValues[len(balanceValues)-3])
}
