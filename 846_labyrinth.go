package main

import "fmt"

const N = 110

var n, m int
var g [N][N]int
var d [N][N]int
var q [N * N]map[int]int

func main() {
	fmt.Scanf("%d%d", &n, &m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", g[i][j])
		}
	}
	bfs()
}

func bfs() {
	hh := 0
	tt := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d[i][j] = -1
		}
	}

	d[0][0] = 0

	for hh <= tt {
		hh++
		t := q[hh]
		for i := 0; i < 4; i++ {

		}
	}

}
