package linkproblems

import (
	"fmt"
	"sync"

	datastruc "github.com/mjk222/demo/algorithm/data_struc"
)

func PrintListReversingly(n *datastruc.Node) {
	stack := datastruc.NewStack()
	if n != nil {
		fmt.Println("info")
		for n.Data != nil {
			stack.Push(n.Data)
			if n.Next == nil {
				break
			}
			n = n.Next
		}
	}
	for i := stack.Length; i > 0; i-- {
		element, err := stack.Pop()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Printf("%v  <-  ", element)
	}

}

var mutex sync.Mutex

func PrintListReversinglyTest() {
	mutex.Lock()
	list := new(datastruc.Node)
	list.TailInterpolation(1)
	list.TailInterpolation(2)
	list.TailInterpolation("majikun")
	mutex.Unlock()
	PrintListReversingly(list)
}
