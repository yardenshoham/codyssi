package main

import (
	"bytes"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func lastNumber(line []byte) int64 {
	parts := bytes.Fields(line)
	num, err := strconv.Atoi(string(parts[len(parts)-1]))
	if err != nil {
		panic(err)
	}
	return int64(num)
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := bytes.Split(input, []byte{'\n'})
	add := big.NewInt(lastNumber(lines[0]))
	multiply := big.NewInt(lastNumber(lines[1]))
	raise := big.NewInt(lastNumber(lines[2]))

	sum := big.NewInt(0)
	for i := 4; i < len(lines); i++ {
		num := lastNumber(lines[i])
		if num%2 == 0 {
			sum.Add(sum, big.NewInt(num))
		}
	}
	sum.Exp(sum, raise, big.NewInt(0))
	sum.Mul(sum, multiply)
	sum.Add(sum, add)
	fmt.Println(sum.String())
}
