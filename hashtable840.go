package main

import "fmt"

const N = 100003

var h [N]int
var e [N]int
var ne [N]int
var idx int

func main() {
	for i := 0; i < N; i++ {
		h[i] = -1
	}
	var n int
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		x := 0
		opt := ""
		fmt.Scanf("%s%d", &opt, &x)
		if opt == "I" {
			insert(x)
		} else {
			res := find(x)
			fmt.Println(res)
		}
	}

}

func insert(x int) {
	k := (x%N + N) % N
	e[idx] = x
	ne[idx] = h[k]
	h[k] = idx
	idx++
}

func find(x int) string {
	k := (x%N + N) % N
	for i := h[k]; i != -1; i = ne[i] {
		if e[i] == x {
			return "Yes"
		}
	}
	return "No"
}
func getPrimeNum(x int) int {
	var primeNum int
	for i := x; ; i++ {
		flag := true
		for j := 2; j*j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeNum = i
			break
		}
	}
	return primeNum
}
