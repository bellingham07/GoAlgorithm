package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	s := "ADOBECODEBANC"
	t1 := "ABC"
	//window := minWindow(s, t1)
	window := minWindow2(s, t1)
	fmt.Println(window)
}
