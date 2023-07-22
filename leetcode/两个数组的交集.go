package leetcode

//https://leetcode.cn/problems/intersection-of-two-arrays/

func intersection(nums1 []int, nums2 []int) []int {
	m1 := map[int]int{}
	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if _, ok := m1[nums1[i]]; !ok {
			m1[nums1[i]]++
		}
	}
	for i := 0; i < len(nums2); i++ {
		if m1[nums2[i]] > 0 {
			res = append(res, nums2[i])
			m1[nums2[i]]--
		}
	}
	return res
}
