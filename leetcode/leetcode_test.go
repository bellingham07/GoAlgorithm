package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	//nums := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	nums := []int{1, 2, 1}
	arrayLen := totalFruit(nums)
	fmt.Println(arrayLen)
}
