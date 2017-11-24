package main

import (
	"fmt"
)

func main() {
	var sum int64 = 2
	for i, j, curr := 1, 2, 3; curr <= 4000000; i, j = j, curr {
		curr = i + j
		if curr%2 == 0 {
			sum += int64(curr)
		}
	}
	fmt.Println(sum)
}
