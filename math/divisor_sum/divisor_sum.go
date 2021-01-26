package main

import "fmt"

func main() {

	n := 0
	fmt.Scanln(&n)
	dmap := make(map[int]int)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scanln(&x)
		for i := 2; i <= x/i; i++ {
			if x%i == 0 {
				for x%i == 0 {
					x /= i
					dmap[i]++
				}
			}
		}
		if x > 1 {
			dmap[x]++
		}
	}

	mod := int(1e9 + 7)
	res := 1
	for key, value := range dmap {
		sum := 1
		for ; value > 0; value-- {
			sum = (sum*key + 1)
			sum = sum % mod
		}
		res = res * sum % mod
	}
	fmt.Println(res)

}
