package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	var k int
	fmt.Scanf("%d%d", &n, &k)
	q := make([]int, n+1)

	//sanner read n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	a := make([]int, 0)
	for scanner.Scan() {
		if scanner.Text() != "" {
			val, _ := strconv.Atoi(scanner.Text())
			a = append(a, val)
		}
	}
	if scanner.Err() != nil {
		fmt.Printf("error: %s\n", scanner.Err())
	}

	hh := 0
	tt := -1
	for i := 0; i < n; i++ {
		for hh <= tt && q[hh] < i-k+1 {
			hh++
		}
		for hh <= tt && a[q[tt]] >= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Printf("%d ", a[q[hh]])
		}
	}
	fmt.Println("")

	hh = 0
	tt = -1
	for i := 0; i < n; i++ {
		for hh <= tt && q[hh] < i-k+1 {
			hh++
		}
		for hh <= tt && a[q[tt]] <= a[i] {
			tt--
		}
		tt++
		q[tt] = i
		if i >= k-1 {
			fmt.Printf("%d ", a[q[hh]])
		}
	}
}
