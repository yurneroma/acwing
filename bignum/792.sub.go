package main

import "fmt"

func main() {
	var a, b string
	fmt.Scanf("%s", &a)
	fmt.Scanf("%s", &b)
	sliceA := make([]int, 0)
	sliceB := make([]int, 0)
	for i := len(a) - 1; i >= 0; i-- {
		sliceA = append(sliceA, int(rune(a[i]-'0')))
	}

	for i := len(b) - 1; i >= 0; i-- {
		sliceB = append(sliceB, int(rune(b[i]-'0')))
	}

	if cmp(sliceA, sliceB) {
		res := sub(sliceA, sliceB)
		for i := len(res) - 1; i >= 0; i-- {
			fmt.Printf("%d", res[i])
		}
	} else {
		res := sub(sliceB, sliceA)
		fmt.Printf("%s", "-")
		for i := len(res) - 1; i >= 0; i-- {
			fmt.Printf("%d", res[i])
		}
	}

}

//compare the size of a and b
func cmp(a []int, b []int) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}

	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true

}

//sub  a.size > b.size
func sub(a []int, b []int) []int {
	res := make([]int, 0)
	t := 0
	for i := 0; i < len(a); i++ {
		t = a[i] - t
		if i < len(b) {
			t -= b[i]
		}
		res = append(res, (t+10)%10)
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	res = trimZero(res)
	return res
}

func trimZero(a []int) []int {
	for len(a) > 1 && a[len(a)-1] == 0 {
		a = a[:len(a)-1]
	}
	return a
}
