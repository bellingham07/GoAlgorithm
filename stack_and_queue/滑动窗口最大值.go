package stack_and_queue

import (
	"container/heap"
	"sort"
)

var a []int

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return a[h.IntSlice[i]] > a[h.IntSlice[j]] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow2(nums []int, k int) []int {
	a = nums
	q := &hp{make([]int, k)}
	for i := 0; i < k; i++ {
		q.IntSlice[i] = i
	}
	heap.Init(q)

	n := len(nums)
	ans := make([]int, 1, n-k+1)
	ans[0] = nums[q.IntSlice[0]]
	for i := k; i < n; i++ {
		heap.Push(q, i)
		for q.IntSlice[0] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, nums[q.IntSlice[0]])
	}
	return ans
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	stack := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		stack = append(stack, nums[i])
		if len(stack) == k {
			res = append(res, max(stack))
			stack = stack[1:]
		}
	}
	return res
}

func max(nums []int) int {
	m := nums[0]
	for i := 0; i < len(nums); i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	return m
}
