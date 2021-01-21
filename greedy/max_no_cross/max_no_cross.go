package main

import (
	"fmt"
	"sort"
)

func main() {
	rangeList := make([]Range, 0)
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		a, b := 0, 0
		fmt.Scanln(&a, &b)
		elem := Range{l: a, r: b}
		rangeList = append(rangeList, elem)
	}

	sort.SliceStable(rangeList, func(i, j int) bool {
		return rangeList[i].r < rangeList[j].r
	})

	res := 0
	ed := int(-2e9)
	for i := 0; i < n; i++ {
		if rangeList[i].l > ed {
			res++
			ed = rangeList[i].r
		}
	}
	fmt.Println(res)

}

type Range struct {
	l int
	r int
}
