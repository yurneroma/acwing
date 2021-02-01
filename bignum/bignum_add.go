package main

import "fmt"

const N = 1e6 + 10

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)

	va := make([]int, 0)
	vb := make([]int, 0)

	for i := len(a) - 1; i >= 0; i-- {
		va = append(va, int(a[i])-'0')
	}

	for i := len(b) - 1; i >= 0; i-- {
		vb = append(vb, int(b[i])-'0')
	}

	res := bigAdd(va, vb)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}

func bigAdd(a, b []int) []int {
	res := make([]int, 0)
	t := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		if i < len(a) {
			t += a[i]
		}
		if i < len(b) {
			t += b[i]
		}
		res = append(res, t%10)
		t /= 10
	}

	if t != 0 {
		res = append(res, t)
	}
	return res
}
