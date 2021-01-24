package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//合并果子
//典型的haffman树问题

func main() {
	pq := make(PriorQueue, 0)
	heap.Init(&pq)

	n := 0
	fmt.Scanln(&n)
	rd := bufio.NewReader(os.Stdin)
	intArr := readline(rd)
	for i := 0; i < n; i++ {
		heap.Push(&pq, intArr[i])
	}

	var res int

	for len(pq) > 1 {
		a := heap.Pop(&pq)
		b := heap.Pop(&pq)
		res += a.(int) + b.(int)
		heap.Push(&pq, a.(int)+b.(int))
	}

	fmt.Println(res)
}

type PriorQueue []int

func (pq PriorQueue) Less(i, j int) bool { return pq[i] < pq[j] }

func (pq PriorQueue) Len() int { return len(pq) }

func (pq PriorQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorQueue) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *PriorQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func readline(rd *bufio.Reader) []int {
	intArr := make([]int, 0)
	line, _ := rd.ReadString('\n')
	params := strings.Fields(line)
	for _, elem := range params {
		intelem, _ := strconv.Atoi(elem)
		intArr = append(intArr, intelem)
	}
	return intArr
}
