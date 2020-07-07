package main

import "fmt"

const N = 5

var size int
var heap [N]int

func main() {
	var n int
	var opt string
	var k, x int
	fmt.Scanf("%d", &n)

	for ; n > 0; n-- {
		fmt.Scanf("%s%d%d", &opt, &k, &x)
		if opt == "I" {
			size++
			fmt.Println("x is: ", k)
			heap[size] = k
			heap[1] = heap[size]
			up(1)
		}
		if opt == "PM" {
			fmt.Println(heap[1])
		}
		if opt == "DM" {
			heap[1] = heap[size]
			size--
			down(1)
		}
		if opt == "D" {

		}
		if opt == "C" {

		}

	}
}

func up(i int) {
	for i/2 > 0 && heap[i] < heap[i/2] {
		heap[i], heap[i/2] = heap[i/2], heap[i]
		i = i >> 1
	}
}

func down(i int) {
	t := i
	if heap[2*i] < heap[i] && 2*i <= size {
		t = 2 * i
	}
	if heap[t] > heap[2*i+1] && 2*i+1 <= size {
		t = 2*i + 1
	}
	heap[i], heap[t] = heap[t], heap[i]
	down(t)
}
