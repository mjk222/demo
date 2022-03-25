package datastruc

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

// 赋值Node
func assignNode() *Node {
	head := new(Node)
	head.Data = 0
	n1 := new(Node)
	head.Next = n1
	n1.Data = 1
	return head
}

// 链表长度
func (n *Node) LengthOfNode() int {
	if n != nil {
		count := 1
		for n.Next != nil {
			count++
		}
		return count
	}
	return 0
}

// 遍历链表输出内容
func ShowNode(n *Node) {
	for n != nil {
		fmt.Printf("%s  ->  ", n.Data)
		n = n.Next
	}
	fmt.Printf("\n")
}

// 尾插法
func (pHead *Node) TailInterpolation(value interface{}) error {
	// fmt.Println("The value before modification is ")
	// ShowNode(pHead)

	if pHead == nil {
		return fmt.Errorf("the stack is nil")
	}
	list := pHead
	node := &Node{
		Data: value,
		Next: nil,
	}
	if list.Data == nil {
		list.Data = value
		return nil
	}
	for list.Next != nil {
		list = list.Next
	}
	list.Next = node

	// fmt.Println("The value after modification is ")
	// ShowNode(pHead)
	return nil
}

// 头插法
func HeadInterpolation() {
	pHead := assignNode()
	fmt.Println("The value before modification is ")
	ShowNode(pHead)

	if pHead == nil {
		pHead = new(Node)
	}

	list := pHead
	head := new(Node)
	head.Data = 48
	head.Next = list
	pHead = head

	fmt.Println("The value after modification is ")
	ShowNode(pHead)
}

// 获取并删除头
func (pHead *Node) PopHead() (interface{}, error) {
	if pHead == nil {
		return nil, fmt.Errorf("this is an empty stack")
	}
	if pHead.Next == nil {
		return pHead.Data, nil
	}
	head := pHead
	pHead = head.Next
	ShowNode(pHead)
	return head.Data, nil
}
