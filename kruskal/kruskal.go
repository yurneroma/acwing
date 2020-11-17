package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//INF
const (
	INF = 0x3f3f3f3f
	N   = 100010
	M   = 200010
)

var (
	n, m   int
	p      [N]int
	reader = bufio.NewReader(os.Stdin)
)

//Edge matrix
type Edge struct {
	A int
	B int
	W int
}

//Kruskal function
func Kruskal(edges []*Edge) int {
	sort.SliceStable(edges, func(i, j int) bool {
		return edges[i].W < edges[j].W
	})

	for i := 1; i <= n; i++ {
		p[i] = i
	}

	res, cnt := 0, 0
	for i := 0; i < m; i++ {
		a := edges[i].A
		b := edges[i].B
		w := edges[i].W
		a = find(a)
		b = find(b)
		if a != b {
			p[a] = b
			res += w
			cnt++
		}
	}

	if cnt < n-1 {
		return INF
	}

	return res
}

//并查集， 找到祖宗节点返回
func find(x int) int {
	if p[x] != x {
		p[x] = find(p[x]) //路径压缩， x 不是祖宗节点，接着往上找， 找到之后，把p[x] 指向祖宗节点， 并返回
	}
	return p[x]
}

func main() {
	fmt.Fscan(reader, &n, &m)

	edges := make([]*Edge, 0)
	for i := 0; i < m; i++ {
		a, b, w := 0, 0, 0
		fmt.Fscan(reader, &a, &b, &w)
		edges = append(edges, &Edge{a, b, w})
	}

	t := Kruskal(edges)
	if t == INF {
		fmt.Println("impossible")
	} else {
		fmt.Println(t)
	}
}
