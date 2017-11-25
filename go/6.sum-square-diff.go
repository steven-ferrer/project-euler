package main

import (
	"fmt"
)

func squaresSum(n int64) int64 {
	if n == 0 {
		return 0
	}
	return (n * n) + squaresSum(n-1)
}

func naturalSum(n int64) int64 {
	if n == 0 {
		return 0
	}
	return n + naturalSum(n-1)
}

func main() {
	var n int64 = 100
	sqSum := squaresSum(n)
	natSum := naturalSum(n)
	natSumSq := natSum * natSum
	diff := natSumSq - sqSum
	fmt.Println(diff)
}
