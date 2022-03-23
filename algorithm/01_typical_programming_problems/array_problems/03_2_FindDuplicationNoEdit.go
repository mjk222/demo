package arrayproblems

import "fmt"

// 面试题3（二）：不修改数组找出重复的数字
// 题目：在一个长度为n+1的数组里的所有数字都在1到n的范围内，所以数组中至
// 少有一个数字是重复的。请找出数组中任意一个重复的数字，但不能修改输入的
// 数组。例如，如果输入长度为8的数组{2, 3, 5, 4, 3, 2, 6, 7}，那么对应的
// 输出是重复的数字2或者3。

func (d *Data) getDuplication(length int) (int, error) {
	if d.Nums2 == nil || length == 0 {
		return -1, fmt.Errorf("invalid nums")
	}
	start := 1
	end := length - 1
	for end >= start {
		middle := start + ((end - start) >> 1)
		count := countRange(d.Nums2, length, start, middle)
		fmt.Printf("start is %v,middle is %v,end is %v\n", start, middle, end)
		fmt.Println("count is ", count)
		switch {
		case end == start:
			if count > 1 {
				return start, nil
			}
			return -1, fmt.Errorf("nums have no duplication")
		case count > middle-start+1:
			end = middle
		case count <= middle-start+1:
			start = middle + 1
		default:
			return -1, fmt.Errorf("switch err")
		}
	}
	return -1, nil
}

func countRange(nums []int, length int, start int, end int) int {
	count := 0
	for i := 0; i < length; i++ {
		if nums[i] >= start && nums[i] <= end {
			count++
		}
	}
	return count
}
