package main

import "fmt"

func main() {
	cnt := 1
	divisorMap := make(map[int]int)
	n := 0
	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		x := 0
		fmt.Scanln(&x)
		divisorNum(x, divisorMap)
	}
	mod := int(1e9 + 7)
	for _, value := range divisorMap {
		cnt *= (value + 1)
		cnt = cnt % mod
	}
	fmt.Println(cnt)
}

func divisorNum(n int, divisorMap map[int]int) {
	for i := 2; i <= n/i; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
				divisorMap[i]++
			}
		}
	}

	if n > 1 {
		divisorMap[n]++
	}
}
