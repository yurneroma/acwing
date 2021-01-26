package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scanln(&x)
		divisors := getDivisors(x)
		for i := 0; i < len(divisors); i++ {
			fmt.Printf("%d ", divisors[i])
		}
		fmt.Println("")
	}
}

//试除法求约数， 从 1 到  n/i ， 如果能被整除， 说明是约数。
func getDivisors(n int) []int {
	divisors := make([]int, 0)
	for i := 1; i <= n/i; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}
	return divisors
}
