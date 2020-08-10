package main

import (
	"fmt"
)

/*
*dijkstra  稀疏图， 使用堆优化
 */
func main() {

	//init the h[N]
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	fmt.Scanf("%d%d", &n, &m)
	for ; m > 0; m-- {
		a, b, c := 0, 0, 0
		fmt.Scanf("%d%d%d", &a, &b, &c)
		add(a, b, c)
	}

	val := dijkstra()
	fmt.Println(val)

}

const N = 100010

var e [N]int
var ne [N]int
var h [N]int
var idx int
var n, m int
var dist [N]int
var w [N]int

func add(a, b, c int) {
	e[idx] = b
	ne[idx] = h[a]
	h[a] = idx
	idx++
	w[idx] = c
}

type item struct {
	Num int
	Dis int
}

var st [N]bool

func (heap []*item) pop() {

}
func (heap []*item) top() {

}

func (heap []*item) push() {

}

func (heap []*item) update() {

}

//dijkstra
func dijkstra() int {
	//init the dist
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}
	dist[1] = 0

	heap := make([]*item, 0)
	heap = append(heap, &item{Num: 1, Dis: 0})

	//
	for len(heap) > 0 {
		t := heap.top()
		heap.pop()
		ver := t.Num
		dis := t.Dis
		if st[ver] {
			continue
		}
		st[ver] = true

		//
		for i := h[ver]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[ver]+w[i] {
				dist[j] = dist[ver] + w[i]
				heap.push(&item{Num: j, Dis: dist[j]})
			}
		}

	}

	if dist[n] == 0x3f3f3f3f {
		return -1
	}
	return dist[n]

}
