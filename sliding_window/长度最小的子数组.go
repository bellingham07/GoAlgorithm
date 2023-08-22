package sliding_window

// https://leetcode.cn/problems/minimum-size-subarray-sum/

//输入：target = 7, nums = [2,3,1,2,4,3]
//输出：2
//解释：子数组 [4,3] 是该条件下的长度最小的子数组。

// minSubArrayLen 暴力法 但不管用 leetcode会超时
func minSubArrayLen(target int, nums []int) int {
	res := len(nums) + 1
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				sub := j - i + 1
				if sub < res {
					res = sub
				}
				break
			}
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// 滑动窗口
// 确定终止位置 如何去移动起始位置
// 起始位置i 终止位置j
// 什么时候移动起始位置 在这个题里面就是i-j这个集合里面的值 >= target

func minSubArrayLen2(target int, nums []int) int {
	i := 0               // 起始位置
	res := len(nums) + 1 // 开局设定 最大值 方便后续赋值
	sum := 0
	// 动态移动结束位置
	for j := 0; j < len(nums); j++ {
		sum += nums[j]      // 记录当前区间的值
		for sum >= target { // 相当于 while true
			subL := j - i + 1 // 计算当前长度
			if subL < res {   // 比已知长度小 重新赋值
				res = subL
			}
			// 滑动窗口 移动起始位置 来不断缩小区间
			sum -= nums[i]
			i++
		}
	}
	// 如果长度为len(nums)+1 说明没找到 直接返回0
	if res == len(nums)+1 {
		return 0
	}
	return res
}
