package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 1010

var f [N]int

func main() {
	n := 0
	fmt.Scanln(&n)
	rd := bufio.NewReader(os.Stdin)
	intArr := readline(rd)
	maxNum := 0
	//force search
	for i := 0; i < len(intArr); i++ {
		f[i] = 1
		for j := 0; j < i; j++ {
			if intArr[j] < intArr[i] {
				f[i] = max(f[i], f[j]+1)
			}
		}
		maxNum = max(maxNum, f[i])
	}
	fmt.Println(maxNum)

}

func readline(rd *bufio.Reader) []int {
	line, _ := rd.ReadString('\n')
	params := strings.Fields(line)
	intArr := make([]int, 0)
	for _, elem := range params {
		intElem, _ := strconv.Atoi(elem)
		intArr = append(intArr, intElem)
	}
	return intArr
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
