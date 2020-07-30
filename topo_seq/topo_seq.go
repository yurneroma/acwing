package main

import "fmt"

const N = 100010

var e [N]int
var ne [N]int
var h [N]int
var idx int
var invec [N]int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

var n, m int

var queue [N]int

func main() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}

	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < m; i++ {
		a, b := 0, 0
		fmt.Scanf("%d%d", &a, &b)
		add(a, b)
		invec[b]++
	}

	if !toposort() {
		fmt.Println(-1)
	} else {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", queue[i])
		}
		fmt.Println("")
	}

}

func toposort() bool {
	hh, tt := 0, -1
	for i := 1; i <= n; i++ {
		if invec[i] == 0 {
			tt++
			queue[tt] = i
		}
	}
	for hh <= tt {
		val := queue[hh]
		hh++
		for i := h[val]; i != -1; i = ne[i] {
			elem := e[i]
			invec[elem]--
			if invec[elem] == 0 {
				tt++
				queue[tt] = elem
			}
		}
	}

	return tt == n-1
}
