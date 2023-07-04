package sort

import (
	"fmt"
	"testing"
)

var nums = []int{9, 5, 3, 6, 8, 1, 2, 5, 4, 7, 8}

func TestSort(t *testing.T) {
	fmt.Println("bubble sort")
	BubbleSort(nums)
	fmt.Println("select sort")
	SelectSort(nums)
	fmt.Println("insert sort")
	InsertSort(nums)
}
