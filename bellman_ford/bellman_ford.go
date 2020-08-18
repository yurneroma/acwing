package main

import "fmt"

var n, m, k int

func main() {
	fmt.Scanf("%d%d%d", &n, &m, &k)
	for i := 0; i < m; i++ {
		a, b, c := 0, 0, 0
		fmt.Scanf("%d%d%d", &a, &b, &c)
		edges[i] = edge{a, b, c}
	}
	bellman_ford()

}

const N = 510
const M = 10010

type edge struct {
	a int
	b int
	w int
}

var dist [N]int
var edges [M]edge
var last [N]int

func bellman_ford() {
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	dist[1] = 0
	for i := 0; i < k; i++ {
		copy(last[0:], dist[0:])
		for i := 0; i < m; i++ {
			e := edges[i]
			dist[e.b] = min(dist[e.b], last[e.a]+e.w)
		}
	}

	if dist[n] > 0x3f3f3f3f/2 {
		fmt.Println("impossible")
	} else {
		fmt.Println(dist[n])
	}

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
