package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
[//]:golang 二维数组解法。
状态方程： f[i][j] 表示前i个物品，体积<=j时，最大值。
当背包装不下第i个物品时：  f[i][j] = f[i-1][j]
当背包装的下时,有两种情况:
	1. 选i，  f[i][j] = f[i][j-v[i]] + w[i]
	2. 不选i， f[i][j] = f[i-1][j]
	f[i][j] = max(f[i-1][j-v[i]] + w[i], f[i-1][j])

*/

const (
	N = 1010
	M = 1010
)

var f [N][M]int
var v [N]int
var w [N]int

func main() {
	rd := bufio.NewReader(os.Stdin)
	in := readline(rd)
	n, m := in[0], in[1]

	for i := 1; i <= n; i++ {
		in := readline(rd)
		v[i], w[i] = in[0], in[1]
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			f[i][j] = f[i-1][j]
			if j >= v[i] {
				f[i][j] = max(f[i][j], f[i-1][j-v[i]]+w[i])
			}
		}
	}
	fmt.Println(f[n][m])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func readline(rd *bufio.Reader) []int {
	line, _ := rd.ReadString('\n')
	params := strings.Fields(line)
	res := make([]int, len(params))
	for i, param := range params {
		res[i], _ = strconv.Atoi(param)
	}
	return res
}
