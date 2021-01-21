package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 100010

var f [N]int

func main() {
	n := 0
	fmt.Scanln(&n)
	res := 0
	r := bufio.NewReader(os.Stdin)
	intArr := readline(r)

	// 使用dp， 复杂度O(n^2), 数据规模10w的话，会超时。
	//solution1
	for i := 0; i < len(intArr); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if intArr[j] < intArr[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		res = max(res, f[i])
	}
	fmt.Println(res)

	//solution2

}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func readline(r *bufio.Reader) []int {
	line, _ := r.ReadString('\n')
	params := strings.Fields(line)
	intArr := make([]int, 0)
	for _, elem := range params {
		intElem, _ := strconv.Atoi(elem)
		intArr = append(intArr, intElem)
	}
	return intArr
}
