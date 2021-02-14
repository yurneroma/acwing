package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)

	//check whether the number is exist
	checkMap := make(map[int]int)
	s := make([]int, n+10)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}

	res := 0

	j := 0
	for i := 0; i < n; i++ {
		checkMap[s[i]]++
		for j <= i && checkMap[s[i]] > 1 {
			checkMap[s[j]]--
			j++
		}
		res = max(res, i-j+1)
	}
	fmt.Println(res)

}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
