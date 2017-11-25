package main

import (
	"fmt"
)

func isPrime(n int64) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	var i int64 = 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}

func main() {
	var n int64 = 10001

	//var primes []int64

	var k int64 = 0
	var cnt int64 = 1
	for k < n {
		if isPrime(cnt) {
			k++
		}
		cnt++
	}

	fmt.Println(cnt - 1)
}
