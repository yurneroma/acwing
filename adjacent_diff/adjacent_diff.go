package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, m := 0, 0
	fmt.Scanln(&n, &m)
	q := make([]int, n+10)
	b := make([]int, n+10)

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 100000*16) //输入范围最大十万个数， 每个整数占8个字节。 申请 10w * 16 防止溢出。
	scanner.Buffer(buf, len(buf))

	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	for i := 1; i <= n; i++ {
		q[i], _ = strconv.Atoi(line[i-1])
		b[i] = q[i] - q[i-1]
	}

	var questionParam [3]int
	for m > 0 {
		m--
		scanner.Scan()
		temp := strings.Split(scanner.Text(), " ")
		for k, v := range temp {
			questionParam[k], _ = strconv.Atoi(v)
		}
		l, r, c := questionParam[0], questionParam[1], questionParam[2]

		b[l] += c
		b[r+1] -= c
	}

	for i := 1; i <= n; i++ {
		q[i] = q[i-1] + b[i]
		fmt.Printf("%d ", q[i])
	}

}
