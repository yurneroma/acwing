package main

import "fmt"

/*思想
1. 确定分界点
2. 调整区间，   l <= x,   x <= r
3. 递归处理左右两边,  两边有序了， 整体就排序完成了。
*/

const N = 100010

var q [N]int
var n int

func main() {
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&q[i])
	}

	quickSort(&q, 0, n-1)
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}

func quickSort(q *[N]int, l, r int) {
	if l >= r {
		return
	}

	x := q[(l+r)>>1]
	i := l - 1
	j := r + 1

	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}
		for {
			j--
			if q[j] <= x {
				break
			}
		}
		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	quickSort(q, l, j)
	quickSort(q, j+1, r)

}
