package main

import "fmt"

//N is the num of verticle
const N = 510

const INF = 0x3f3f3f3f

//const INF = 1<<63 - 1
//const INF = 0x7FF0000000000000

var g [N][N]int
var dist [N]int
var st [N]bool
var n, m int

func main() {
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g[i][j] = INF
		}
	}
	x, y, z := 0, 0, 0
	for i := 0; i < m; i++ {
		fmt.Scanf("%d%d%d", &x, &y, &z)
		g[x][y] = min(g[x][y], z)
	}

	dis := dijkstra()
	fmt.Println(dis)
}

func min(a int, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

func dijkstra() int {
	for i := 0; i < N; i++ {
		dist[i] = INF
	}
	dist[1] = 0

	for i := 0; i < n; i++ {
		t := -1

		//找到当前未标记过，且离源点最近的点
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}
		}

		st[t] = true

		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], dist[t]+g[t][j])
		}
	}
	if dist[n] == INF {
		return -1
	}
	return dist[n]
}
