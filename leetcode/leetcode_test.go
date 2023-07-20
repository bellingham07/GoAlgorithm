package leetcode

import (
	"testing"
)

func TestLeetCode(t *testing.T) {
	obj := Constructor()
	param_1 := obj.Get(5)
	obj.AddAtHead(1)
	obj.AddAtTail(6)
	obj.AddAtIndex(3, 5)
	obj.DeleteAtIndex(3)
}
