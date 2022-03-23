package notes

import "fmt"

func PointNotes() bool {
	var i1 = 5
	i2 := 2
	fmt.Println(&i1)
	intP := &i1
	fmt.Println(intP, *intP)
	if i1 == i2 {
		return true
	} else {
		return false
	}
}

// 问题：列举Go语言中*号的所有用法
