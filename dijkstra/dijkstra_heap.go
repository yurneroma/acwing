package main

import (
	"fmt"
)

const N = 100010

var h [N]int
var e [N]int
var ne [N]int
var w [N]int
var idx int

//邻接表存边和权重
func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++

}

var n, m int

func main() {
	// fmt.Scanf("%d%d", &n, &m)
	// for i := 0; i < N; i++ {
	// 	h[i] = -1
	// }

	// //邻接表存储图
	// for ; m > 0; m-- {
	// 	a, b, c := 0, 0, 0
	// 	fmt.Scanf("%d%d%d", &a, &b, &c)
	// 	add(a, b, c)
	// }
	dijkstra()
}

var dist [N]int
var st [N]int

type Item struct {
	Ver int
	Dis int
}

type Heap []*Item

func (heap Heap) Len() int { return len(heap) }

func (heap *Heap) Init() {
	n := heap.Len()
	for i := n/2 - 1; i >= 0; i-- {
		heap.Down(i)
	}

}

func (heap *Heap) Push(item *Item) {
	*heap = append(*heap, item)
}

func (heap *Heap) Up() {

}

func (heap *Heap) Swap(i, j int) {
	(*heap)[i], (*heap)[j] = (*heap)[j], (*heap)[i]
}

func (heap *Heap) Down(idx int) {
	t := idx
	n := len(*heap)
	if 2*idx+1 < n && (*heap)[idx*2+1].Dis < (*heap)[t].Dis {
		t = 2*idx + 1
	}
	if 2*idx+2 < n && (*heap)[2*idx+2].Dis < (*heap)[t].Dis {
		t = 2*idx + 2
	}

	if t != idx {
		heap.Swap(idx, t)
		heap.Down(t)
	}

}

func (heap *Heap) Pop() {

}
func (heap *Heap) Top() {

}

func (heap *Heap) Update() {

}

func (heap *Heap) Print() {
	n := len(*heap)
	for i := 0; i < n; i++ {
		fmt.Println((*heap)[i])
	}
}
func dijkstra() int {
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
	}

	dist[1] = 0

	item := &Item{
		Ver: 1,
		Dis: 0,
	}

	heap := make(Heap, 0)
	heap.Push(item)
	heap.Init()

	for len(heap) > 0 {
		t := heap.Top()
		heap.Pop()

		ver := t.Ver
		dis := t.Dis
		if st[ver] {
			continue
		}
		st[ver] = true

		for i := h[ver]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[ver]+w[i] {
				dist[j] = dist[ver] + w[i]
				heap.Push(&Item{Ver: j, Dis: dist[j]})
			}
		}
	}

	if dist[n] == 0x3f3f3f3f {
		return -1
	}
	return dist[n]

}
