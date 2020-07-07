package main

import "fmt"

const N = 100010

var p [N]int
var size [N]int

func main() {
	opt := ""
	n, m := 0, 0
	fmt.Scanf("%d%d", &n, &m)
	for i := 0; i < n; i++ {
		p[i] = i
		size[i] = 1
	}
	for ; m > 0; m-- {
		a, b := 0, 0
		fmt.Scanf("%s%d%d", &opt, &a, &b)
		if opt == "C" {
			if find(a) != find(b) {
				size[find(b)] += size[find(a)]
				p[find(a)] = find(b)
			}
		}
		if opt == "Q1" {
			if find(a) == find(b) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
		if opt == "Q2" {
			fmt.Println(size[a])
		}
	}
}

func find(x int) int {
	if p[x] != x {
		p[x] = find(x)
	}
	return p[x]
}
