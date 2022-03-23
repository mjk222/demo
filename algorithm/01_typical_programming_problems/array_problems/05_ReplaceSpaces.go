package arrayproblems

import "fmt"

// 面试题5：替换空格
// 题目：请实现一个函数，把字符串中的每个空格替换成"%20"。例如输入“We are happy.”，
// 则输出“We%20are%20happy.”。
func ReplaceBlank() {
	d := Data{
		Str1: " We are  happy. ",
	}
	s := []byte(d.Str1)
	p1 := len(s) - 1
	count := 0
	for _, v := range s {
		if v == 32 {
			count += 2
		}
	}
	s1 := make([]byte, count)
	s = append(s, s1...)
	p2 := len(s) - 1

	for ; p1 >= 0; p1-- {
		if s[p1] != 32 {
			fmt.Printf("p1 is %v,p2 is %v\n", p1, p2)
			s[p2] = s[p1]
			p2--
			continue
		}
		s[p2] = 48
		s[p2-1] = 50
		s[p2-2] = 37
		p2 -= 3
	}
	fmt.Println("Str1 is ", string(s))
}
