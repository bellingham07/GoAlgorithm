package leetcode

import (
	"fmt"
	"testing"
)

func TestLeetCode(t *testing.T) {
	s := "adc"
	s2 := "dcda"
	//window := minWindow(s, t1)
	window := checkInclusion(s, s2)
	fmt.Println(window)
}
