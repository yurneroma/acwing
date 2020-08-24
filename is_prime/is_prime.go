package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		x := 0
		fmt.Scanf("%d", &x)
		res := is_prime(x)
		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}

func is_prime(x int) bool {
	if x < 2 {
		return false
	}

	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true

}
