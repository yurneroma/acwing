package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		a, b := 0, 0
		fmt.Scanln(&a, &b)
		fmt.Println(gcd(a, b))
	}
}

// 欧几里得算法，   d | a,   d | b  ==>  d | ax + by
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
