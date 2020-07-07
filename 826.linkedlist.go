package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func main() {
	m := 0
	fmt.Scanf("%d", &m)

	lk := &LinkedList{&Node{Val: 0, Next: nil}}
	for ; m > 0; m-- {
		var t string
		var p, q int
		fmt.Scanf("%s%d%d", &t, &p, &q)
		switch t {
		case "H":
			HeadInsert(lk, p)
		case "I":
			Insert(lk, p, q)
		case "D":
			Del(lk, p)
		}

	}

	PrintLk(lk)

}

func PrintLk(lk *LinkedList) {
	p := lk.Head
	for p != nil {
		fmt.Printf("%d ", p.Val)
		if p.Next != nil {
			p = p.Next
		}
	}
}

//HeadInsert  insert x after the head
func HeadInsert(lk *LinkedList, x int) {
	head := lk.Head
	node := &Node{Val: x, Next: nil}
	node.Next = head.Next
	head.Next = node
}

//Insert after the kth insert with value x
func Insert(lk *LinkedList, k, x int) {

}

//Del the kth elem
func Del(lk *LinkedList, k int) {

}
