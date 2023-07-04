package search

import (
	"fmt"
	"testing"
)

func TestBinary(t *testing.T) {
	nums := []int{12, 23, 34, 45, 56, 67, 78, 89, 99}
	//insert := searchInsert(nums, 55)
	binary := Binary(55, nums)
	fmt.Println(binary)
	if binary != -1 {
		fmt.Println(nums[binary])
	}

	//fmt.Println(insert)
	//fmt.Println(nums[insert])
}
