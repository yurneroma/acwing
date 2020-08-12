package main

import (
	"container/heap"
	"fmt"
)

const N = 100010

var e [N]int
var ne [N]int
var w [N]int
var h [N]int
var dist [N]int
var st [N]bool

var idx int

var n, m int

func main() {
	fmt.Scanf("%d%d", &n, &m)

	//init
	for i := 0; i < N; i++ {
		dist[i] = 0x3f3f3f3f
		h[i] = -1
	}
	dist[1] = 0

	for ; m > 0; m-- {
		a, b, c := 0, 0, 0
		fmt.Scanf("%d%d%d", &a, &b, &c)
		add(a, b, c)
	}

	val := dijkstra()
	fmt.Println(val)
}

func dijkstra() int {
	start := &Item{Ver: 1, Dis: 0}
	pq := make(Pq, 0)
	heap.Push(&pq, start)

	for pq.Len() > 0 {
		t := heap.Pop(&pq).(*Item)
		ver := t.Ver
		if st[ver] {
			continue
		}
		st[ver] = true

		for i := h[ver]; i != -1; i = ne[i] {
			j := e[i]
			if dist[j] > dist[ver]+w[i] {
				dist[j] = dist[ver] + w[i]
				heap.Push(&pq, &Item{Ver: j, Dis: dist[j]})
			}
		}
	}

	if dist[n] == 0x3f3f3f3f {
		return -1
	}
	return dist[n]
}

func add(a, b, c int) {
	e[idx] = b
	w[idx] = c
	ne[idx] = h[a]
	h[a] = idx
	idx++
}

type Item struct {
	Ver int
	Dis int
}

type Pq []*Item

func (pq Pq) Len() int {
	return len(pq)
}

func (pq Pq) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq Pq) Less(i, j int) bool {
	return pq[i].Dis < pq[j].Dis
}

func (pq *Pq) Pop() interface{} {
	n := pq.Len()
	old := *pq
	item := old[n-1]
	old[n-1] = nil

	*pq = old[0 : n-1]
	return item

}

func (pq *Pq) Push(item interface{}) {
	*pq = append(*pq, item.(*Item))
}
