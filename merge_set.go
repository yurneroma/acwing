package main

import "fmt"

const N = 100010
var p [N]int
func main() {
	n, m:= 0,0
	fmt.Scanf("%d%d",&n,&m)
	for i:=0;i<n;i++{
		p[i] = i
	}
	for ;m>0;m--{
		opt := ""
		a, b := 0,0
		fmt.Scanf("%s%d%d",&opt,&a,&b)
		switch opt {
		case "M":
			p[find(a)] = find(b)
		case "Q":
			if find(a) == find(b) {
				fmt.Println("Yes")
			}else{
				fmt.Println("No")
			}

		}
	}
}

func find(x int)int{
	if p[x] != x{
		p[x] = find(p[x])
	}
	return p[x]
}



