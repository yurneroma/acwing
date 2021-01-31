package main

import "fmt"

const N = 100010

var q [N]int

func main() {
	n, k := 0, 0
	fmt.Scanln(&n, &k)
	for i := 0; i < n; i++ {
		fmt.Scan(&q[i])
	}

	quickSort(&q, 0, n-1)
	fmt.Println(q[k-1])
}

func quickSort(q *[N]int, l, r int) {
	if l >= r {
		return
	}

	x := q[(l+r)>>1]
	i := l - 1
	j := r + 1

	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}

		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	quickSort(q, l, j)
	quickSort(q, j+1, r)

}
