package main

import (
	"fmt"
)

func main() {
	min := 1
	max := 20

	n := 1

	found := false
	for !found {
		curr := true
		for i := min; i <= max; i++ {
			curr = curr && (n%i == 0)
		}
		if curr {
			break
		}
		n++
	}

	fmt.Println(n)
}
