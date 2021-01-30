package main

import (
	"fmt"
	"math"
)

const N = 110

var n int
var a [N][N]float64

func main() {
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ {
			fmt.Scan(&a[i][j])
		}
	}

	t := gauss()
	if t == 0 {
		for i := 0; i < n; i++ {
			fmt.Printf("%.2f\n", a[i][n])
		}
	} else if t == 1 {
		fmt.Println("Infinite group solutions")
	} else {
		fmt.Println("No solution")
	}

}

const eps = 1e-6

func gauss() int {
	var c, r int
	for ; c < n; c++ {
		t := r
		for i := r; i < n; i++ {
			if math.Abs(a[i][c]) > math.Abs(a[t][c]) {
				t = c
			}
		}
		if math.Abs(a[t][c]) < eps {
			continue
		}

		//swap
		for i := c; i <= n; i++ {
			a[t][i], a[r][i] = a[r][i], a[t][i]
		}

		//convert first num of this line to 1
		for i := n; i >= c; i-- {
			a[r][i] /= a[r][c]
		}

		for i := r + 1; i < n; i++ {
			if math.Abs(a[i][c]) > eps {
				for j := n; j >= c; j-- {
					a[i][j] -= a[r][j] * a[i][c]
				}
			}
		}
		r++
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n+1; j++ {
			a[i][n] -= a[i][j] * a[j][n]
		}
	}
	if r < n {
		for i := r; i < n; i++ {
			if math.Abs(a[i][n]) > eps {
				return 2 //no solution
			}
		}
		return 1 //Infinite solutions
	}

	return 0
}
