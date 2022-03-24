package arrayproblems

import "fmt"

// 面试题4：二维数组中的查找
// 题目：在一个二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按
// 照从上到下递增的顺序排序。请完成一个函数，输入这样的一个二维数组和一个
// 整数，判断数组中是否含有该整数。
func Find() bool {
	num := 6
	d := Data{
		Matrix: [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}},
	}
	if d.Matrix == nil || len(d.Matrix) == 0 {
		return false
	}
	col := len(d.Matrix) - 1
	row := 0
	fmt.Println(col, row)
	for col >= 0 && row < len(d.Matrix[0]) {
		fmt.Println("row is ", row)
		fmt.Println("col is ", col)
		fmt.Println("num is ", d.Matrix[row][col])
		if d.Matrix[row][col] != num {
			if num < d.Matrix[row][col] {
				col--
			} else {
				row++
			}
		} else {
			return true
		}
	}
	return false
}
