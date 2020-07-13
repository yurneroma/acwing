package main

import "fmt"

const N = 9

var path [N][N]string
var dg [N]bool
var udg [N]bool
var col [N]bool
var n int

func main() {
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			path[i][j] = "."
		}
	}
	dfs(0)
}

func dfs(u int) {
	if u == n {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Printf("%s ", path[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
		return
	}

	for i := 0; i < n; i++ {
		if col[i] == false && dg[u+i] == false && udg[n+u-i] == false {
			path[u][i] = "Q"
			col[i] = true
			dg[u+i] = true
			udg[n+u-i] = true
			dfs(u + 1)
			col[i] = false
			dg[u+i] = false
			udg[n+u-i] = false
			path[u][i] = "."

		}
	}
}
