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

func exponentialisPecunia(add, multiply, raise *big.Int, num int64) *big.Int {
	bigNum := big.NewInt(num)
	bigNum.Exp(bigNum, raise, big.NewInt(0))
	bigNum.Mul(bigNum, multiply)
	bigNum.Add(bigNum, add)
	return bigNum
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
	threshold := big.NewInt(15000000000000)
	best := big.NewInt(0)
	for i := 4; i < len(lines); i++ {
		num := lastNumber(lines[i])
		bigNum := big.NewInt(num)
		if exponentialisPecunia(add, multiply, raise, num).Cmp(threshold) <= 0 && bigNum.Cmp(best) == 1 {
			best = bigNum
		}
	}
	fmt.Println(best.String())
}
