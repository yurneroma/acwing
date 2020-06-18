package main

import "fmt"

const N = 100010

func main() {
	n := 0
	fmt.Scanf("%d", &n)

	var tt = 0
	stack := make([]int, N)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)
		for stack[tt] >= x && tt > 0 {
			tt--
		}
		if tt > 0 {
			fmt.Printf("%d ", stack[tt])
		} else {
			fmt.Printf("%d ", -1)
		}
		tt++
		stack[tt] = x
	}

}
