package interview

import (
	"fmt"
	"testing"
)

func TestInterview(t *testing.T) {

	a1 := []int{4, 9}
	//a1 := []int{9, 9, 9, 9}
	a2 := []int{1}
	ints := add(a1, a2)
	fmt.Println(ints)
}
