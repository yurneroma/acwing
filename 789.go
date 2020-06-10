package main

import "fmt"

func main() {

	n, m := 0, 0
	fmt.Scanf("%d%d", &n, &m)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}

	for m > 0 {
		m--

		x := 0
		fmt.Scanf("%d", &x)
		l := 0
		r := n - 1
		for l < r {
			mid := (l + r) >> 1
			if q[mid] < x {
				l = mid + 1
			} else {
				r = mid
			}
		}

		if q[l] != x {
			fmt.Println(-1, -1)
		} else {
			l1 := l
			r1 := n - 1
			for l1 < r1 {
				mid := (l1 + r1 + 1) >> 1
				if q[mid] <= x {
					l1 = mid
				} else {
					r1 = mid - 1
				}
			}
			fmt.Println(l, l1)
		}

	}
	return
}
