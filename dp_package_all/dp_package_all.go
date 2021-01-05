package main

import "fmt"

const N = 1010

var res [N]int
var v [N]int
var w [N]int

func main() {

	var n, m int
	fmt.Scanln(&m, &n)
	for i := 1; i <= n; i++ {
		a, b := 0, 0
		fmt.Scanln(&a, &b)
		v[i], w[i] = a, b
	}

	for i := 1; i <= n; i++ {
		for j := v[i]; j <= m; j++ {
			res[j] = max(res[j], res[j-v[i]]+w[i])
		}
	}

	fmt.Println(res[m])

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
