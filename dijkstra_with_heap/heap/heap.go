package piorqueue

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
