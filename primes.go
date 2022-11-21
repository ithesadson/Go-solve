//prime numbers from 1 to n
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var n int = int(math.Pow(10, 2))
	var primes []int

	for i := 2; i < n+1; i++ {
		primes = append(primes, i)
	}
	//fmt.Println(primes)
	for x := 0; x < int(n/2); x++ {
		if primes[x] != 0 {
			for i := x + primes[x]; i < n-1; i += primes[x] {
				primes[i] = 0
			}
		}
	}
	sort.Ints(primes)

	var zero int = 0
	for j := 0; j < len(primes); j++ {
		if primes[j] == 0 {
			zero += 1
		}
	}
	primes = primes[zero:]
	fmt.Println(primes)
}
