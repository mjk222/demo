package datastruc

import "fmt"

// 定义栈
type Stack struct {
	Top    *Node
	Length int
}

// 新建
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// 入栈
func (s *Stack) Push(value interface{}) error {
	if s == nil {
		return fmt.Errorf("the stack is nil")
	}
	node := new(Node)
	node.Data = value
	node.Next = s.Top
	s.Top = node
	s.Length++
	return nil
}

// 出栈
func (s *Stack) Pop() (data interface{}, err error) {
	if s.Length == 0 {
		return nil, fmt.Errorf("the stack is nil")
	}
	data = s.Top.Data
	s.Top = s.Top.Next
	s.Length--
	return data, nil
}

// 测试函数
func StackTest() {
	stack := NewStack()
	//s1:=new(Stack)
	//stack.Push(1)
	//stack.Push("majikun")
	fmt.Println(stack.Top)
	fmt.Println(stack.Length)
	result, err := stack.Pop()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("result is ", result)
	fmt.Println(stack.Top)
	fmt.Println(stack.Length)
}
