package main

import "fmt"

/*

题目描述:
给定一个n个点m条边的有向图，图中可能存在重边和自环。

所有边的长度都是1，点的编号为1~n。

请你求出1号点到n号点的最短距离，如果从1号点无法走到n号点，输出-1。

输入格式
第一行包含两个整数n和m。

接下来m行，每行包含两个整数a和b，表示存在一条从a走到b的长度为1的边。

输出格式
输出一个整数，表示1号点到n号点的最短距离。

数据范围
1≤n,m≤105
输入样例：
4 5
1 2
2 3
3 4
1 3
1 4
输出样例：
1

*/

var n int
var dis = make(map[int]int)

var queue = make([]int, 0)

func bfs() int {
	dis[1] = 0
	queue = append(queue, 1)
	for len(queue) > 0 {
		t := queue[0]
		fmt.Println("cur t is ", t)
		queue = queue[1:]
		for i := h[t]; i != -1; i = ne[i] {
			val := e[i]
			fmt.Println("val ", val)
			if dis[val] == -1 {
				dis[val] = dis[t] + 1
				fmt.Println("dis val ", dis[val])
				queue = append(queue, val)
			}
		}
	}
	return dis[n]
}

const N = 110

var e [N]int
var ne [N * 2]int
var h [N * 2]int
var idx int

func add(a, b int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

func main() {
	m := 0
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		dis[i] = -1
		h[i] = -1
	}
	fmt.Println(n, m)
	for i := 0; i < m; i++ {
		a, b := 0, 0
		fmt.Scanf("%d%d", &a, &b)
		add(a, b)
	}

	val := bfs()
	fmt.Println(val)

}
