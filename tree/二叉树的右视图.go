package tree

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 0)
	dfs(root, 0, &res)
	return res
}

func dfs(node *TreeNode, deep int, res *[]int) {
	if deep == len(*res) {
		*res = append(*res, node.Val)
	}
	if node.Right != nil {
		dfs(node.Right, deep+1, res)
	}
	if node.Left != nil {
		dfs(node.Left, deep+1, res)
	}
}
