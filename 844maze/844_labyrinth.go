package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const N = 110

var n, m int
var d [N][N]int

/*

0 1 0 0 0
0 1 0 1 0
0 0 0 0 0
0 1 1 1 0
0 0 0 1 0

*/

func str2intList(line string) []int {
	intList := make([]int, 0)
	vals := strings.Split(line, " ")
	for _, val := range vals {
		intval, _ := strconv.Atoi(val)
		intList = append(intList, intval)
	}
	return intList
}
func main() {

	fmt.Scanln(&n, &m)
	g := make([][]int, n)
	scn := bufio.NewScanner(os.Stdin)
	i := 0
	for scn.Scan() {
		line := scn.Text()
		row := str2intList(line)
		g[i] = row
		i++
		if i >= n {
			break
		}
	}

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	minpath := bfs(g)
	fmt.Println(minpath)
}

type Point struct {
	X int
	Y int
}

func bfs(g [][]int) int {
	//d store the value from (0,0) to (i,j)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			d[i][j] = -1
		}
	}
	d[0][0] = 0

	var dx = []int{-1, 0, 1, 0}
	var dy = []int{0, -1, 0, 1}

	hh := 0
	tt := 0
	queue := make([]Point, N*N)
	queue[hh] = Point{X: 0, Y: 0}
	for hh <= tt {
		elem := queue[hh]
		hh++
		for i := 0; i < 4; i++ {
			x := elem.X + dx[i]
			y := elem.Y + dy[i]
			if x >= 0 && y >= 0 && x < n && y < m && g[x][y] == 0 && d[x][y] == -1 {
				d[x][y] = d[elem.X][elem.Y] + 1
				tt++
				queue[tt] = Point{X: x, Y: y}
			}
		}
	}
	return d[n-1][m-1]
}
