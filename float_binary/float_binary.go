package main

import "fmt"

func main() {
	var l, r, mid, x float64
	l, r = -10000, 10000

	fmt.Scanln(&x)

	for r-l > 1e-8 {
		mid = (l + r) / 2
		if mid*mid*mid >= x {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Printf("%.6f", mid)
}
