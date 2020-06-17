package main

import "fmt"

const N = 100010

var l []int = make([]int, N)
var r []int = make([]int, N)
var e []int = make([]int, N)
var idx int

func main() {
	l[1] = 0
	r[0] = 1
	idx = 2
	m := 0
	fmt.Scanf("%d", &m)
	for ; m > 0; m-- {
		var op string
		var k, x int
		fmt.Scanf("%s", &op)
		if op == "L" {
			fmt.Scanf("%d", &x)
			add(0, x)
		} else if op == "R" {
			fmt.Scanf("%d", &x)
			add(l[1], x)
		} else if op == "D" {
			fmt.Scanf("%d", &k)
			del(k + 1)

		} else if op == "IL" {
			fmt.Scanf("%d%d", &k, &x)
			add(l[k+1], x)
		} else {
			fmt.Scanf("%d%d", &k, &x)
			add(k+1, x)

		}
	}
	for i := r[0]; i != 1; i = r[i] {
		fmt.Printf("%d ", e[i])
	}
}

//k的右边插入x
func add(k, x int) {
	e[idx] = x
	r[idx] = r[k]
	l[idx] = k
	l[r[k]] = idx
	r[k] = idx
	idx++
}

func del(k int) {
	l[r[k]] = l[k]
	r[l[k]] = r[k]
}
