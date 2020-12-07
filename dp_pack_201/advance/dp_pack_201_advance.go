package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	N = 1010
	M = 1010
)

var f [N]int
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
		for j := m; j >= v[i]; j-- {
			f[j] = max(f[j], f[j-v[i]]+w[i])
		}
	}
	fmt.Println(f[m])
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
