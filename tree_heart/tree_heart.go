package main

import "fmt"

var n int

func main() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	fmt.Scanf("%d", &n)
	for i := 0; i < n-1; i++ {
		a, b := 0, 0
		fmt.Scanf("%d%d", &a, &b)
		add(a, b)
		add(b, a)
	}

	dfs(1)
	fmt.Println(heart)
}

const N = 100010

var h [N]int
var ne [2 * N]int
var e [2 * N]int
var idx int
var st [N]bool
var ans int = N

//将节点 b 加入到节点a 上, 单向
func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

var heart int = N

func dfs(u int) int {
	st[u] = true
	sum := 1
	size := 0
	for i := h[u]; i != -1; i = ne[i] {
		val := e[i]
		if st[val] {
			continue
		}
		son := dfs(val)
		size = max(son, size)
		sum += son
	}

	size = max(size, n-sum)
	heart = min(heart, size)

	return sum
}

func max(a, b int) int {
	if a <= b {
		return b
	} else {
		return a
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}

}
