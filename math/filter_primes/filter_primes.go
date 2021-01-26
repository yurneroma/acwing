package main

import "fmt"

const N = 1000010

var st [N]bool
var primes [N]int
var cnt int

func main() {
	n := 0
	fmt.Scanln(&n)
	for i := 2; i <= n; i++ {
		if !st[i] {
			primes[cnt] = i
			cnt++
		}

		for j := 0; primes[j] <= n/i; j++ {
			st[primes[j]*i] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
	fmt.Println(cnt)
}
