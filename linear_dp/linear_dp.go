package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 510

var matrix [N][N]int
var f [N][N]int

func main() {
	rd := bufio.NewReader(os.Stdin)
	n := 0
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		intArr := readline(rd)
		for j := 1; j <= i; j++ {
			matrix[i][j] = intArr[j-1]
		}
	}

	for i := n; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			f[i][j] = max(f[i+1][j], f[i+1][j+1]) + matrix[i][j]
		}
	}

	fmt.Println(f[1][1])

}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func readline(rd *bufio.Reader) []int {
	line, _ := rd.ReadString('\n')
	intArr := make([]int, 0)
	strArr := strings.Fields(line)
	for i := 0; i < len(strArr); i++ {
		elem, err := strconv.Atoi(strArr[i])
		if err != nil {
			fmt.Println(err)
		}
		intArr = append(intArr, elem)
	}
	return intArr
}
