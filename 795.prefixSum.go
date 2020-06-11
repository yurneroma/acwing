package main

import "fmt"

func main() {
	n, m := 0, 0
	fmt.Scanf("%d%d", &n, &m)
	q := make([]int, n+1)
	presum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%d", &q[i])
		presum[i] = presum[i-1] + q[i]
	}
	for m > 0 {
		l, r := 0, 0
		fmt.Scanf("%d%d", &l, &r)
		if l > r {
			return
		}
		if r > n || l < 1 {
			return
		}
		sum := presum[r] - presum[l-1]
		fmt.Println(sum)
		m--
	}
}
