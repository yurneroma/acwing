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
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < N; i++ {
		h[i] = -1
	}

	//邻接表存储图
	for ; m > 0; m-- {
		a, b, c := 0, 0, 0
		fmt.Scanf("%d%d%d", &a, &b, &c)
		add(a, b, c)
	}
	val := dijkstra()
	fmt.Println(val)
}

var dist [N]int
var st [N]bool

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
	heap.Up()
}

func (heap *Heap) Up() {
	n := heap.Len()
	for i := n; i > 1; i = i / 2 {
		if (*heap)[i-1].Dis < (*heap)[i/2-1].Dis {
			heap.Swap((i - 1), (i/2 - 1))
		} else {
			break
		}
	}
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
	old := heap
	n := heap.Len()
	(*heap)[0] = (*heap)[n-1]
	(*old)[n-1] = nil
	*heap = (*old)[:n-1]
	heap.Down(0)
}

func (heap *Heap) Top() *Item {
	if heap.Len() > 0 {
		return (*heap)[0]
	} else {
		return nil
	}
}

func (heap *Heap) Print() {
	n := len(*heap)
	for i := 0; i < n; i++ {
		fmt.Println((*heap)[i])
	}
}
func dijkstra() int {
	for i := 0; i < N; i++ {
		dist[i] = 0x3f
	}

	dist[1] = 0

	item := &Item{
		Ver: 1,
		Dis: 0,
	}

	heap := make(Heap, 0)
	heap.Push(item)

	for len(heap) > 0 {
		t := heap.Top()
		heap.Pop()

		ver := t.Ver
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
