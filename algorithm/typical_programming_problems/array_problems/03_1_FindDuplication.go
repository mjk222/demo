package arrayproblems

import (
	"fmt"
)

// 面试题3（一）：找出数组中重复的数字
// 题目：在一个长度为n的数组里的所有数字都在0到n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
// 也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。例如，如果输入长度为7的数组{2, 3, 1, 0, 2, 5, 3}，
// 那么对应的输出是重复的数字2或者3。

//方法①：排序数组，顺序扫描

//方法②：顺序扫描数组，哈希表里没有则加入哈希表，有则找到一个重复的数字

//方法③：当前数字和数组下标不同时，交换当前数字对应下标的数字，直到当前数字和数组下标相同
//注意：数组为空，数字小于0或大于等于数组长度
func RepeatNumberThree() {
	data := new(Data)
	data.Nums1 = []int{2, 1, 3, 5, 4}
	data.Nums2 = []int{1, 2, 6, 4, 5, 3}
	result1, err := data.duplicate(len(data.Nums1))
	if err != nil {
		fmt.Println(err)
	}
	result2, err := data.getDuplication(len(data.Nums2))
	fmt.Println("duplicate num is ", result2)
	if result1 {
		fmt.Printf("There is at least one duplicate number is %v", data.Duplication)
		return
	}
	fmt.Println(err)
}

func (d *Data) duplicate(length int) (bool, error) {
	if d.Nums1 == nil || length <= 0 {
		return false, fmt.Errorf("nums is nil")
	}
	for i := 0; i < length; i++ {
		if d.Nums1[i] < 0 || d.Nums1[i] >= length {
			return false, fmt.Errorf("nums[%d] err", i)
		}
	}
	for index := 0; index < length; index++ {
		for i := index; i < length; i++ {
			if d.Nums1[i] == i {
				break
			}
			if d.Nums1[i] == d.Nums1[d.Nums1[i]] {
				d.Duplication = d.Nums1[i]
				return true, nil
			}
			d.Nums1[i], d.Nums1[d.Nums1[i]] = d.Nums1[d.Nums1[i]], d.Nums1[i]
		}
	}
	return false, fmt.Errorf("there is no duplication number")
}
