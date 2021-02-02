package main

import "fmt"

func main() {
	n, m := 0, 0
	fmt.Scanln(&n, &m)
	q := make([]int, n+1)
	partialSum := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&q[i])
		partialSum[i] += partialSum[i-1] + q[i]
	}

	for ; m > 0; m-- {
		l, r := 0, 0
		fmt.Scanln(&l, &r)
		fmt.Println(partialSum[r] - partialSum[l-1])
	}

}
