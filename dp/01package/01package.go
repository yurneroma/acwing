package main

import (
	"bufio"
	"fmt"
	"os"
)

const N = 10

var w [N]int
var v [N]int
var res [N]int

func main() {
	var n, m int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &n, &m)

	for i := 1; i <= n; i++ {
		fmt.Fscan(reader, &v[i], &w[i])
	}

	//res表示体积为j的情况下，  权重的最大值
	for i := 1; i <= n; i++ {
		for j := m; j >= v[i]; j-- {
			res[j] = max(res[j-v[i]]+w[i], res[j])
		}
	}

	fmt.Println(res[m])
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
