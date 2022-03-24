package datastruc

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func assignNode() *Node {
	head := new(Node)
	head.Data = 0
	n1 := new(Node)
	head.Next = n1
	n1.Data = 1
	return head
}

func ShowNode() {
	n := assignNode()
	for n != nil {
		fmt.Println(n)
		n = n.Next
	}
}
