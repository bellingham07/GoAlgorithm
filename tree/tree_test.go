package tree

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {

	/*
	       1
	      / \
	     2   3
	    / \
	   4   5
	*/
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	fmt.Println(averageOfLevels(root))
}
