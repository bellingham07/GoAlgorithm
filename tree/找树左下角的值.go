package tree

func findBottomLeftValue(root *TreeNode) int {
	res := 0
	d := -1
	var ddfs func(node *TreeNode, deep int)
	ddfs = func(node *TreeNode, deep int) {
		if node == nil {
			return
		}
		if deep > d {
			res = node.Val
			d = deep
		}
		ddfs(node.Left, deep+1)
		ddfs(node.Right, deep+1)
	}
	ddfs(root, 0)
	return res
}
