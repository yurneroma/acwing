package main

import "fmt"

const N = 1010

func main() {
	n, m, q := 0, 0, 0
	fmt.Scanln(&n, &m, &q)
	var matrix [N][N]int
	var mpSum [N][N]int

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Scan(&matrix[i][j])
			mpSum[i][j] = mpSum[i-1][j] + mpSum[i][j-1] - mpSum[i-1][j-1] + matrix[i][j]
		}
	}

	for ; q > 0; q-- {
		x1, y1, x2, y2 := 0, 0, 0, 0
		fmt.Scanln(&x1, &y1, &x2, &y2)
		res := mpSum[x2][y2] - mpSum[x2][y1-1] - mpSum[x1-1][y2] + mpSum[x1-1][y1-1]
		fmt.Println(res)

	}
}
