package main

import "fmt"

const N = 100010

var h [N]int
var size int

func main() {
	n, m := 0, 0
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", h[i])
	}
	size = n
	//建堆
	for i := n / 2; i > 0; i-- {
		down(i)
	}

	for ; m > 0; m-- {
		fmt.Printf("%d", h[1])
		h[1] = h[size]
		size--
		down(1)
	}
}

func down(x int) {
	t := x
	if x*2 <= size && h[x*2] < h[t] {
		t = x * 2
	}
	if x*2+1 <= size && h[x*2+1] < h[t] {
		t = 2*x + 1
	}
	if x != t {
		h[x], h[t] = h[t], h[x]
	}
	down(t)
}
