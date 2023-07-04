package sort

import (
	"testing"
)

var nums = []int{9, 5, 3, 6, 8, 1, 2, 5, 4, 7, 8}

func TestSort(t *testing.T) {
	//BubbleSort(nums)
	//SelectSort(nums)
	InsertSort(nums)
}
