package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const N = 1010

var f [N][N]int

func main() {
	m, n := 0, 0
	fmt.Scanln(&m, &n)
	a, b := readInput()
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 0 || j == 0 {
				f[i][j] = 0
			}
			if a[i] == b[j] {
				f[i][j] = max(f[i-1][j-1]+1, f[i][j])
			} else {
				f[i][j] = max(f[i][j-1], f[i-1][j])
			}
		}
	}

	fmt.Println(f[m][n])
}

func readInput() (a, b string) {
	reader := bufio.NewReader(os.Stdin)
	line1, _ := reader.ReadString('\n')
	line2, _ := reader.ReadString('\n')
	a = strings.Split(line1, "\n")[0]
	a = fmt.Sprint(" ", a)
	b = strings.Split(line2, "\n")[0]
	b = fmt.Sprint(" ", b)
	return

}

func max(a, b int) int {
	if a <= b {
		return b
	}

	return a
}
