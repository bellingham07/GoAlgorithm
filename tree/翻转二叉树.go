package tree

// bfs
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
	}

	return root
}

// dfs
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)

	return root
}
