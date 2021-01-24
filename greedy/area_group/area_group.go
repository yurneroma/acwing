package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type Range struct {
	l int
	r int
}

func main() {
	rangeList := make([]*Range, 0)
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		l, r := 0, 0
		fmt.Scanln(&l, &r)
		elem := &Range{l: l, r: r}
		rangeList = append(rangeList, elem)
	}

	sort.SliceStable(rangeList, func(i, j int) bool {
		return rangeList[i].l < rangeList[j].l
	})

	pq := make(IntHeap, 0)
	heap.Init(&pq)
	for i := 0; i < n; i++ {
		if len(pq) == 0 || pq[0] >= rangeList[i].l {
			heap.Push(&pq, rangeList[i].r)
		} else {
			heap.Pop(&pq)
			heap.Push(&pq, rangeList[i].r)
		}
	}
	fmt.Println(pq.Len())
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x

}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
