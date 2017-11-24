package main

import (
	"fmt"
)

func main() {
	sum := 0
	upper := 1000
	for i := 3; i < upper; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
