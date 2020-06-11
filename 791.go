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
	newSlice := make([]int, 0)
	if len(sliceA) < len(sliceB) {
		return add(sliceB, sliceA)
	}
	t := 0
	for i := 0; i < len(sliceA); i++ {
		t += sliceA[i]
		if i < len(sliceB) {
			t += sliceB[i]
		}
		newSlice = append(newSlice, t%10)
		t /= 10
	}
	if t > 0 {
		newSlice = append(newSlice, 1)
	}
	return newSlice
}
