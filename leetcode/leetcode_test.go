package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	s := "acdcaeccde"
	p := "c"
	//window := minWindow(s, t1)
	window := findAnagrams(s, p)
	fmt.Println(window)
}
