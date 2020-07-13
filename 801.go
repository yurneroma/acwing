package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	for i := 0; i < n; i++ {
		val := 0
		for j := 0; j < 64; j++ {
			if q[i]&1 == 1 {
				val++
			}
			q[i] = q[i] >> 1
		}
		fmt.Printf("%d ", val)
	}

}
