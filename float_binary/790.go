package main

import "fmt"

func main() {
	var n float64
	fmt.Scanf("%f", &n)

	var l float64 = -10000
	var r float64 = 10000
	var mid float64

	for r-l > 1e-8 {
		mid = (l + r) / 2
		if mid*mid*mid >= n {
			r = mid
		} else {
			l = mid
		}
	}
	fmt.Printf("%.6f", l)
}
