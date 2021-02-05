package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010

func main() {
	var q [N]int
	var b [N]int

	n, m := 0, 0
	fmt.Scanln(&n, &m)
	reader := bufio.NewReader(os.Stdin)
	line := readline(reader)
	for i := 1; i <= n; i++ {
		q[i] = line[i-1]
		b[i] = q[i] - q[i-1]
	}

	for m > 0 {
		m--
		line = readline(reader)
		l, r, c := line[0], line[1], line[2]
		b[l] += c
		b[r+1] -= c
	}
	for i := 1; i <= n; i++ {
		q[i] = q[i-1] + b[i]
		fmt.Printf("%d ", q[i])
	}
}

func readline(reader *bufio.Reader) []int {
	strline, _ := reader.ReadString('\n')
	params := strings.Fields(strline)
	line := make([]int, len(params))
	for i, elem := range params {
		line[i], _ = strconv.Atoi(elem)
	}
	return line
}
