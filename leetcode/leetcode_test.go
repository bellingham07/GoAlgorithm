package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	arr := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(arr)
	fmt.Println(profit)
}
