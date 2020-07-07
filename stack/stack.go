package main

import "fmt"

const M = 100010

var idx = -1
var stk = make([]int, M)

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	for ; n > 0; n-- {
		opt := ""
		x := 0
		fmt.Scanf("%s", &opt)
		switch opt {
		case "push":
			fmt.Scanf("%d", &x)
			push(x)
		case "query":
			res := query()
			fmt.Println(res)
		case "pop":
			pop()
		case "empty":
			res := empty()
			fmt.Println(res)
		}

	}
}

func push(x int) {
	if idx < M {
		idx++
		stk[idx] = x
	}
}

func query() int {
	if idx < 0 {
		return -1
	}
	return stk[idx]
}

func empty() string {
	if idx < 0 {
		return "YES"
	}
	return "NO"
}

func pop() {
	if idx >= 0 {
		idx--
	}
}
