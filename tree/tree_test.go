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

	p := NewNode(4)
	q := NewNode(5)

	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		}
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		}
		if (root.Val-p.Val)*(root.Val-q.Val) <= 0 {
			return root
		}
	}
	return root
}
