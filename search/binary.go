package search

func Binary(target int, nums []int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		// 防止溢出
		mid := left + ((right - left) / 2)
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	// 没找到 直接返回-1
	return -1
}

// 第一种二分法
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right + 1
}
