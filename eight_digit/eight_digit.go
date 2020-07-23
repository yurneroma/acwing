package main

import (
	"bytes"
	//"bytes"
	"fmt"
	"strings"
)

func main() {
	start := make([]string, 0)
	elem := ""
	for i := 0; i < 9; i++ {
		fmt.Scanf("%s", &elem)
		start = append(start, elem)
	}

	mergeStr := strings.Join(start, "")

	bfs(mergeStr)
}

//bfs + 字符串数组哈希
func bfs(start string) {
	var end string = "12345678x"
	var dx = []int{-1, 0, 1, 0}
	var dy = []int{0, -1, 0, 1}
	dist := make(map[string]int)
	dist[start] = 0
	queue := make([]string, 0)
	queue = append(queue, start)
	for len(queue) != 0 {
		t := queue[0]
		st := strings.Split(t, "")
		index := bytes.IndexByte([]byte(t), 'x')
		x := index % 3
		y := index / 3
		for i := 0; i < 4; i++ {
			tmpx := x + dx[i]
			tmpy := y + dy[i]
			if tmpx >= 0 && tmpy >= 0 && tmpx < 3 && tmpy < 3 {
				tmpIndex := tmpy*3 + tmpx
				//fmt.Println("tmpIndex ", tmpIndex)
				st[index], st[tmpIndex] = st[tmpIndex], st[index]
				//fmt.Println("st is ", st)
				covstr := strings.Join(st, "")
				if dist[covstr] == 0 {
					dist[covstr] = dist[t] + 1
					queue = append(queue, covstr)
				}

				st[index], st[tmpIndex] = st[tmpIndex], st[index]
			}
		}
		queue = queue[1:]
	}
	if dist[end] == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[end])
	}

}
