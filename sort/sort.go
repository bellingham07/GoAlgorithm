package sort

import "fmt"

// SelectSort 选择排序
func SelectSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		min := i
		// 每次循环选出最小的
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		// 将最小的放到最前面
		nums[i], nums[min] = nums[min], nums[i]
	}
	fmt.Println(nums)
}

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		// 每次都能选出当前循环最大的数，所以不需要遍历到最后，
		for j := 1; j < len(nums)-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	fmt.Println("after sort", nums)
}

// InsertSort 插入排序
func InsertSort(nums []int) {
	for i := range nums {
		// 双指针 pre记录前一个位置，动态检验大小
		// current记录当前位置的值
		preIndex := i - 1
		current := nums[i]
		// 第一个条件防止数组越界
		// 第二个条件 则是交换的条件
		for preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			// 滑动数组，因为在插入排序中会出现两个位置的值相同，需要记录下最后一个值的位置
			preIndex--
		}
		// 最后一个值小于current，所以将current放在这个值后面
		nums[preIndex+1] = current
	}
	fmt.Println(nums)
}
