package main

import (
	"fmt"
)

func main() {
	min := 1
	max := 20
	var ns []int64
	for i := min; i <= max; i++ {
		ns = append(ns, int64(i))
	}
	fmt.Println(lcmm(ns...))
}

func lcmm(ns ...int64) int64 {
	if len(ns) == 2 {
		return lcm(ns[0], ns[1])
	}
	newNs := ns[1:]
	return lcm(ns[0], lcmm(newNs...))
}

func gcd(a int64, b int64) int64 {
	var t int64
	for b != 0 {
		t = b
		b, a = a%b, t
	}
	return a
}

func lcm(a int64, b int64) int64 {
	return (a * b / gcd(a, b))
}
