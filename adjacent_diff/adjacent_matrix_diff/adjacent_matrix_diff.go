package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 2000*16)
	scanner.Buffer(buf, len(buf))

	line := readline(scanner)
	n, m, q := line[0], line[1], line[2]

	b := make([][]int, n+10)
	for i := 0; i < len(b); i++ {
		b[i] = make([]int, m+10)
	}

	for i := 1; i <= n; i++ {
		line := readline(scanner)
		for j := 1; j <= m; j++ {
			insert(b, i, j, i, j, line[j-1])
		}
	}

	for ; q > 0; q-- {
		line := readline(scanner)
		insert(b, line[0], line[1], line[2], line[3], line[4])
	}

	for i := 1; i <= n; i++ {
		res := make([]string, m)
		for j := 1; j <= m; j++ {
			b[i][j] += b[i-1][j] + b[i][j-1] - b[i-1][j-1]
			//fmt.Printf("%d ", b[i][j]) 频繁调用fmt会造成超时，使用 strings.Join优化
			res[j-1] = strconv.Itoa(b[i][j])
		}
		fmt.Println(strings.Join(res, " "))
	}

}

func insert(b [][]int, x1, y1, x2, y2, c int) {
	b[x1][y1] += c
	b[x2+1][y1] -= c
	b[x1][y2+1] -= c
	b[x2+1][y2+1] += c
}

func readline(scanner *bufio.Scanner) []int {
	scanner.Scan()
	params := strings.Split(strings.TrimSpace(scanner.Text()), " ")
	line := make([]int, len(params))
	for i, param := range params {
		line[i], _ = strconv.Atoi(param)
	}

	return line
}
