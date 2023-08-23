package tree

import "strconv"

var paths []string

func binaryTreePaths(root *TreeNode) []string {
	paths = []string{}
	find(root, "")
	return paths
}

func find(node *TreeNode, p string) {
	if node != nil {
		p += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			paths = append(paths, p)
		} else {
			p += "->"
			find(node.Left, p)
			find(node.Right, p)
		}
	}

}
