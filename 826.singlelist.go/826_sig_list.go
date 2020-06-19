package main

import "fmt"

const M = 100010

var elem = make([]int, M)
var next = make([]int, M)

//head 头结点下标，  idx: 当前已经用到了哪个点
var head, idx = -1, 0

func main() {
	//init()
	m := 0
	fmt.Scanf("%d", &m)
	for ; m > 0; m-- {
		opt := ""
		k, x := 0, 0
		fmt.Scanf("%s", &opt)
		switch opt {
		case "H":
			fmt.Scanf("%d", &x)
			addToHead(x)
		case "D":
			fmt.Scanf("%d", &k)
			if k == 0 {
				head = next[head]
			} else {
				del(k - 1)
			}
		case "I":
			fmt.Scanf("%d%d", &k, &x)
			add(k-1, x)
		}
	}
	for i := head; i != -1; i = next[i] {
		fmt.Printf("%d ", elem[i])
	}
}

func addToHead(x int) {
	elem[idx] = x
	next[idx] = head
	head = idx
	idx++
}

func add(k, x int) {
	elem[idx] = x
	next[idx] = next[k]
	next[k] = idx
	idx++
}

func del(k int) {
	next[k] = next[next[k]]
}
