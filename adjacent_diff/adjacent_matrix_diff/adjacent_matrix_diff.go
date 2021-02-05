package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1010

func main() {
	n, m, qs := 0, 0, 0
	fmt.Scanln(&n, &m, &qs)
	var q [N][N]int
	var b [N][N]int

	rd := bufio.NewReader(os.Stdin)
	for i := 1; i <= n; i++ {
		line := readline(rd)
		for j := 1; j <= m; j++ {
			q[i][j] = line[j-1]
			b[i][j] = q[i][j] - 
		}
	}

}

func readline(rd *bufio.Reader) []int {
	strline, _ := rd.ReadString('\n')
	params := strings.Fields(strline)
	line := make([]int, len(params))
	for i, elem := range params {
		line[i], _ = strconv.Atoi(elem)
	}
	return line
}
