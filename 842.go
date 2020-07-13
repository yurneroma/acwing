package main

import "fmt"

const N = 10

var n int
var path [N]int
var st [N]bool

func main() {
	fmt.Scanf("%d", &n)
	dfs(0)
}

func dfs(u int) {
	if u == n {
		for i := 1; i <= n; i++ {
			fmt.Printf("%d", path[i])
		}
		fmt.Println("")
		return
	}
	for i := 1; i <= n; i++ {
		if st[i] == false {
			path[u] = i
			st[i] = true
			dfs(u + 1)
			st[i] = false

		}
	}

}
