package main

import "fmt"

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		x := 0
		fmt.Scanf("%d", &x)
		dive_prim(x)
		fmt.Println("")
	}
}

//分解质因数
func dive_prim(x int) {

	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			s := 0
			for x%i == 0 {
				x = x / i
				s++
			}
			fmt.Println(i, s)
		}
	}

	if x > 1 {
		fmt.Println(x, 1)
	}
}
