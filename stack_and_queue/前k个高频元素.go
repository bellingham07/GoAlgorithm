package stack_and_queue

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	res := make([]int, 0)
	for _, v := range m {
		res = append(res, v)
	}

	sort.Ints(res)

	res = res[len(res)-k:]
	temp := make([]int, 0)
	for i := k - 1; i >= 0; i-- {
		if len(temp) == k {
			break
		}
		for index, v := range m {
			if v == res[i] {
				temp = append(temp, index)
				m[index] = -1
			}
		}
	}
	return temp
}
