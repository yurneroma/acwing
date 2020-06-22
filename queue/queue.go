package main

import "fmt"

const M = 100010

var hh int = 0
var tt int = -1
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
	tt++
	stk[tt] = x
}

func pop() {
	if hh <= tt {
		hh++
	}
}

func query() int {
	if hh <= tt {
		return stk[hh]
	}
	return 0
}

func empty() string {
	if hh <= tt {
		return "NO"
	}
	return "YES"
}
