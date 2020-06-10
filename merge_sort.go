package main

import "fmt"

func mergeSort(q, temp []int, l, r int) {
	if l >= r {
		return
	}
	mid := (l + r) >> 1
	mergeSort(q, temp, l, mid)
	mergeSort(q, temp, mid+1, r)
	i := l
	j := mid + 1
	k := 0
	for i <= mid && j <= r {
		if q[i] <= q[j] {
			temp[k] = q[i]
			k++
			i++
		} else {
			temp[k] = q[j]
			k++
			j++
		}
	}
	for i <= mid {
		temp[k] = q[i]
		k++
		i++
	}
	for j <= r {
		temp[k] = q[j]
		k++
		j++
	}

	j = 0
	for i := l; i <= r; i++ {
		q[i] = temp[j]
		j++
	}
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	temp := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	mergeSort(q, temp, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}
