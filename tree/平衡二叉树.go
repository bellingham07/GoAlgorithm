package tree

func isBalanced(root *TreeNode) bool {
	h := getHeight(root)
	if h == -1 {
		return false
	}
	return true
}

// 返回以该节点为根节点的二叉树的高度，如果不是平衡二叉树了则返回-1
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := getHeight(root.Left), getHeight(root.Right)
	if l == -1 || r == -1 {
		return -1
	}
	if l-r > 1 || r-l > 1 {
		return -1
	}
	return max1(l, r) + 1
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
