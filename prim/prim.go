package main

import "fmt"

const N = 510
const INF = 0x3f3f3f3f

var g [N][N]int
var st [N]bool
var dist [N]int
var n, m int

func main() {

	//init graph N*N inf
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			g[i][j] = INF
		}
	}

	fmt.Scanf("%d%d", &n, &m)

	for ; m > 0; m-- {
		a, b, c := 0, 0, 0
		fmt.Scanf("%d%d%d", &a, &b, &c)
		g[a][b] = min(g[a][b], c)
		g[b][a] = min(g[a][b], c)
	}

	t := prim()
	if t == INF {
		fmt.Println("impossible")
	} else {
		fmt.Println(t)
	}

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b

}

func prim() int {
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	res := 0

	for i := 0; i < n; i++ {
		t := -1
		for j := 1; j <= n; j++ {
			if !st[j] && (t == -1 || dist[t] > dist[j]) {
				t = j
			}

		}

		if i > 0 && dist[t] == INF {
			return INF
		}

		st[t] = true
		if i > 0 {
			res += dist[t]
		}

		for j := 1; j <= n; j++ {
			dist[j] = min(dist[j], g[t][j])
		}

	}
	return res

}
