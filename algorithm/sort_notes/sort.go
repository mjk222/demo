package sortnotes

import "fmt"

// 冒泡排序
// ①序列最右端有一个天平，比较两数大小
// ②天平左移，选出最小值
// ③回到①
func BubbleSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		fmt.Printf("第%d次循环\n", i)
		for j := 0; j < len(nums)-i; j++ {
			fmt.Println("  排序 ", j+1)
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// 选择排序
// ①从数列中寻找最小值
// ②将最小值和序列的最左边数字交换，回到①
func SelectionSort(nums []int) []int {
	var n []int
	for pos := 0; pos < len(nums); pos++ {
		temp := nums[pos]
		for num := pos + 1; num < len(nums); num++ {
			if nums[num] < temp {
				temp = nums[num]
			}
		}
		n = append(n, temp)
	}
	return n
}

// 插入排序
// ①选出一个数字放到左边的排序区域
// ②比较大小直到左边的数字比自己小
// ③回到①
func InsertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
	return nums
}

// 堆排序
// ①按照降序建堆
// ②取出根节点放在序列最右边
// ③回到①
func HeapSort() {

}

// 归并排序
// ①分割序列
// ②合并两个相邻的序列，比较两个序列的首个数字
//  小的数字放在序列左边
// ③回到②
func MergeSort() {

}

// 快速排序
