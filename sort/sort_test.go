package sort

import (
	"fmt"
	"testing"
)

var nums = []int{9, 5, 3, 6, 8, 1, 2, 5, 4, 7, 8}

func TestSort(t *testing.T) {
	//a1 := []int{1, 3, 5, 7, 9}
	//a2 := []int{2, 4, 6, 8, 10, 45, 84, 656, 48}
	//bubble(nums)
	//selectS(nums)
	//q1(nums, 0, len(nums)-1)
	//fmt.Println(nums)

	res := m1(nums)
	fmt.Println(res)
}

func m1(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := m1(arr[:mid])
	right := m1(arr[mid:])

	return m2(left, right)
}

func m2(left, right []int) []int {
	i, j := 0, 0
	l, r := len(left), len(right)
	res := []int{}
	for {
		if i >= l || j >= r {
			break
		}
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	if i != l {
		res = append(res, left[i:]...)
	} else {
		res = append(res, right[j:]...)
	}

	return res
}
