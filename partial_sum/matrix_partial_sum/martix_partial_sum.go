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
	n, m, q := 0, 0, 0
	fmt.Scanln(&n, &m, &q)
	var matrix [N][N]int
	var mpSum [N][N]int

	reader := bufio.NewReader(os.Stdin)
	for i := 1; i <= n; i++ {
		line := readline(reader)
		for j := 1; j <= m; j++ {
			matrix[i][j] = line[j-1]
			mpSum[i][j] = mpSum[i-1][j] + mpSum[i][j-1] - mpSum[i-1][j-1] + matrix[i][j]
		}
	}

	for ; q > 0; q-- {
		line := readline(reader)
		x1, y1, x2, y2 := line[0], line[1], line[2], line[3]
		res := mpSum[x2][y2] - mpSum[x2][y1-1] - mpSum[x1-1][y2] + mpSum[x1-1][y1-1]
		fmt.Println(res)

	}
}

func readline(reader *bufio.Reader) []int {
	intArr := make([]int, 0)
	line, _ := reader.ReadString('\n')
	params := strings.Fields(line)
	for _, elem := range params {
		intElem, _ := strconv.Atoi(elem)
		intArr = append(intArr, intElem)
	}
	return intArr
}
