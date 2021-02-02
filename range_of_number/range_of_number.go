package main

import "fmt"

func main() {
	n, a := 0, 0
	fmt.Scan(&n, &a)

	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&q[i])
	}

	for ; a > 0; a-- {
		x := 0
		fmt.Scanln(&x)

		l := 0
		r := n - 1
		mid := 0
		for l < r {
			mid = (l + r) >> 1
			if q[mid] < x {
				l = mid + 1
			} else {
				r = mid
			}
		}

		if q[l] != x {
			fmt.Println("-1 -1")
		} else {
			fmt.Printf("%d ", l)
			l = 0
			r = n - 1
			for l < r {
				mid = (l + r + 1) >> 1
				if q[mid] > x {
					r = mid - 1
				} else {
					l = mid
				}
			}
			fmt.Printf("%d \n", l)

		}

	}
}
