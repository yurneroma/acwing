package main

import "fmt"

const N = 100010

var son [N][26]int
var cnt [N]int
var idx int

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		opt := ""
		a := ""
		fmt.Scanf("%s%s", &opt, &a)
		switch opt {
		case "I":
			insert(a)
		case "Q":
			res := query(a)
			fmt.Println(res)
		}
	}

}

func insert(str string) {
	//p=0 , 根节点
	p := 0
	for i := 0; i < len(str); i++ {
		chi := int(str[i] - 'a')
		if son[p][chi] == 0 {
			idx++
			son[p][chi] = idx
		}
		p = son[p][chi]
	}
	cnt[p]++
}

func query(str string) int {
	p := 0
	for i := 0; i < len(str); i++ {
		chi := int(str[i] - 'a')
		if son[p][chi] == 0 {
			return 0
		}
		p = son[p][chi]
	}
	return cnt[p]
}
