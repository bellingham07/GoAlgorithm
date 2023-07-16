package offer

//https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/

func findRepeatNumber(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return v
		}
		m[v]++
	}
	return 0
}
