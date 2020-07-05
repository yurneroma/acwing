package main

import (
	"fmt"
)

const N = 100010
const P = 131

var h [N]int64
var p [N]int64

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	var s string
	fmt.Scanln(&s)
	p[0] = 1
	for i := 1; i <= n; i++ {
		p[i] = p[i-1] * P
		h[i] = h[i-1]*P + int64(s[i-1])
	}

	for ; m > 0; m-- {
		var l1, r1, l2, r2 int
		fmt.Scanln(&l1, &r1, &l2, &r2)
		if get(l1, r1) == get(l2, r2) {
			fmt.Println("Yes")
		} else {

			fmt.Println("No")
		}
	}

}

func get(l, r int) int64 {
	return h[r] - h[l-1]*p[r-l+1]
}
