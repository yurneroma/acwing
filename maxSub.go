package main

import "fmt"

const N = 100010

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	max := 0
	countArr := make([]int, N)
	j := 0
	for i := 0; i < n; i++ {
		countArr[q[i]]++
		for j <= i && countArr[q[i]] > 1 {
			countArr[q[j]]--
			j++
		}
		max = getMax(max, i-j+1)
	}
	fmt.Println(max)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
