package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&q[i])
	}

	tmp := make([]int, n)
	mergeSort(q, tmp, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}

}

func mergeSort(q, tmp []int, l, r int) {
	if l >= r {
		return
	}

	mid := (l + r) >> 1
	mergeSort(q, tmp, l, mid)
	mergeSort(q, tmp, mid+1, r)

	i := l
	j := mid + 1
	k := 0
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			tmp[k] = q[i]
			i++
			k++
		} else {
			tmp[k] = q[j]
			k++
			j++
		}

	}

	for i <= mid {
		tmp[k] = q[i]
		k++
		i++
	}

	for j <= r {
		tmp[k] = q[j]
		k++
		j++
	}

	j = 0
	i = l
	for i <= r {
		q[i] = tmp[j]
		i++
		j++
	}
}
