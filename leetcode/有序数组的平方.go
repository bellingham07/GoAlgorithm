package leetcode

// https://leetcode.cn/problems/squares-of-a-sorted-array/
func sortedSquares(nums []int) []int {
	left := 0
	right := len(nums) - 1
	k := right
	res := make([]int, len(nums))
	// 使用是双指针，因为给定的是有序的数组，所以最左边和最右边总有一个数是最大的
	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[k] = nums[right] * nums[right]
			k--
			right--
		} else {
			res[k] = nums[left] * nums[left]
			k--
			left++
		}
	}
	return res
}
