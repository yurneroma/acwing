package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	n := 0
	fmt.Scanln(&n)
	rd := bufio.NewReader(os.Stdin)
	intArr := readline(rd)
	sort.SliceStable(intArr, func(i, j int) bool {
		return intArr[i] < intArr[j]
	})

	res := 0
	for i := 0; i < n; i++ {
		res += abs(intArr[i] - intArr[i/2])
	}
	fmt.Println(res)
}

func readline(rd *bufio.Reader) []int {
	intArr := make([]int, 0)
	line, _ := rd.ReadString('\n')
	params := strings.Fields(line)
	for _, elem := range params {
		intElem, _ := strconv.Atoi(elem)
		intArr = append(intArr, intElem)
	}
	return intArr
}

func abs(a int) int {
	if a <= 0 {
		return -a
	}
	return a
}
