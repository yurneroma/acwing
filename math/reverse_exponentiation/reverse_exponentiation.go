package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	for ; n > 0; n-- {
		a, p := 0, 0
		fmt.Scanln(&a, &p)
		res := qmi(a, p-2, p)
		if a%p != 0 {
			fmt.Println(res)
		} else {
			fmt.Println("impossible")
		}
	}

}

func qmi(a, k, p int) int64 {
	var res int64
	res = 1
	for k != 0 {
		if k&1 == 1 {
			res = res * int64(a) % int64(p)
		}
		k >>= 1
		a = a * a % p
	}

	return res

}
