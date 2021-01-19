package main

import "fmt"

const N = 1010

func main() {
	n, m := 0, 0
	var a, b string
	fmt.Scanln(&n)
	fmt.Scanln(&a)
	fmt.Scanln(&m)
	fmt.Scanln(&b)
	var dp [N][N]int
	a = fmt.Sprint(" ", a)
	b = fmt.Sprint(" ", b)

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + 1
			if a[i] == b[j] {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	fmt.Println(dp[n][m])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
