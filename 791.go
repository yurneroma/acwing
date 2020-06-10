package main

import "fmt"

func main() {

	var m, n string
	fmt.Scanf("%s", &m)
	fmt.Scanf("%s", &n)

	if len(m) == 0 || len(n) == 0 {
		return
	}
	sliceA := make([]int, 0)
	sliceB := make([]int, 0)
	for i := len(m) - 1; i >= 0; i-- {
		sliceA = append(sliceA, int(rune(m[i])-'0'))
	}
	for i := len(n) - 1; i >= 0; i-- {
		sliceB = append(sliceB, int(rune(n[i])-'0'))
	}

	res := add(sliceA, sliceB)
	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("%d", res[i])
	}
}

func add(sliceA, sliceB []int) []int {
	i := 0
	newSlice := make([]int, 0)
	agrev := 0
	for i < len(sliceA) && i < len(sliceB) {
		res := sliceA[i] + sliceB[i] + agrev
		val := res % 10
		agrev = res / 10
		newSlice = append(newSlice, val)
		i++
	}
	if i == len(sliceA) && i == len(sliceB) {
		if agrev > 0 {
			newSlice = append(newSlice, 1)
		}
	} else {
		for i < len(sliceA) {
			newSlice = append(newSlice, (sliceA[i] + agrev))
			i++
		}
		for i < len(sliceB) {
			newSlice = append(newSlice, (sliceB[i] + agrev))
			i++
		}
	}

	return newSlice
}
