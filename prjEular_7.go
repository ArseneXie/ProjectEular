package main

import (
	"fmt"
)

func main() {
	//n := 120000
	cnt := 1
	

	primes := make([]int, 1, 20000)
	primes[0] = 2

SIEVE:
	for i := 3; cnt < 10001; i++ {
		for _, p := range primes {
			if i%p == 0 {
				continue SIEVE
			}
		}
		primes = append(primes, i)
		cnt = cnt +1
	}

	fmt.Println(primes[cnt-1])
}