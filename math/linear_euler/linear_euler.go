package main

import "fmt"

const N = 1000010

var st [N]bool
var primes [N]int
var cnt int
var euler [N]int

func main() {
	n := 0
	fmt.Scanln(&n)
	linearEuler(n)
	res := 0
	for i := 1; i <= n; i++ {
		res += euler[i]
	}
	fmt.Println(res)
}

func linearEuler(n int) {
	euler[1] = 1
	for i := 2; i <= n; i++ {
		if !st[i] {
			primes[cnt] = i
			cnt++
			euler[i] = i - 1
		}
		for j := 0; primes[j] <= n/i; j++ {
			t := i * primes[j]
			st[t] = true
			if i%primes[j] == 0 {
				euler[t] = euler[i] * primes[j]
				break
			}
			euler[t] = euler[i] * (primes[j] - 1)
		}
	}
}
