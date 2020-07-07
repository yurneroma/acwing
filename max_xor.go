package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const N = 100010

const M = 3000030

var son [M][2]int
var idx int
var a [N]int

func insert(x int) {
	p := 0
	for i := 30; i >= 0; i-- {
		u := (x >> i) & 1
		if son[p][u] == 0 {
			idx++
			son[p][u] = idx
		}
		p = son[p][u]
	}
}

func query(x int) int {
	p := 0
	val := 0
	for i := 30; i >= 0; i-- {
		u := (x >> i) & 1
		if son[p][u^1] > 0 {
			val = val + 1<<i
			p = son[p][u^1]
		} else {
			p = son[p][u]
		}
	}
	return val

}
func main() {
	n := 0
	fmt.Scanf("%d", &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	a := make([]int, 0)
	for scanner.Scan() {
		if scanner.Text() != "" {
			val, _ := strconv.Atoi(scanner.Text())
			a = append(a, val)
		}
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}

	for i := 0; i < len(a); i++ {
		insert(a[i])
	}
	res := 0
	for i := 0; i < len(a); i++ {
		res = max(res, query(a[i]))
	}

	fmt.Println(res)
}

func max(res, a int) int {
	if res <= a {
		return a
	}
	return res
}
