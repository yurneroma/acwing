package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scanln(&x)

		res := euler(x)
		fmt.Println(res)
	}
}

//容斥原理
//f(n) =  n * (1- 1/p1) * (1- 1/p2) *... * (1- 1/pn)
func euler(n int) int {
	cnt := n
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			cnt = cnt * (i - 1) / i
			for n%i == 0 {
				n /= i
			}
		}
	}

	if n > 1 {
		cnt = cnt * (n - 1) / n
	}
	return cnt
}
