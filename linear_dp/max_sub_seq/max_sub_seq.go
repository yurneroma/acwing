package main

import (
	"bufio"
	"strconv"
	"strings"
)

func main() {

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
