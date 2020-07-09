package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	dfs(0, n)

}

const N = 10

var path [N]int
var st [N]bool

//dfs implements
func dfs(u, n int) {
	if u == n {
		for i := 0; i < n; i++ {
			fmt.Printf("%d ", path[i])
		}
		fmt.Println("")
		return
	}
	for i := 1; i <= n; i++ {
		if st[i] == false {
			path[u] = i
			st[i] = true
			dfs(u+1, n)
			st[i] = false
		}
	}
}
