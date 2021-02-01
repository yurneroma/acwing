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
	res := mergeSort(q, tmp, 0, n-1)
	fmt.Println(res)

}

func mergeSort(q, tmp []int, l, r int) int64 {
	if l >= r {
		return 0
	}

	mid := (l + r) >> 1

	var res int64
	res += mergeSort(q, tmp, l, mid)
	res += mergeSort(q, tmp, mid+1, r)

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
			res += int64(mid - i + 1)
		}
	}

	for i <= mid {
		tmp[k] = q[i]
		i++
		k++
	}
	for j <= r {
		tmp[k] = q[j]
		j++
		k++
	}

	j = 0
	for i := l; i <= r; i++ {
		q[i] = tmp[j]
		j++
	}

	return res
}
