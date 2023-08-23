package tree

// bfs
func sumOfLeftLeaves(root *TreeNode) int {
	queue := []*TreeNode{root}
	sum := 0

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				if node.Left.Left == nil && node.Left.Right == nil {
					sum += node.Left.Val
				}
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
	}

	return sum
}

// dfs
func sumOfLeftLeaves1(root *TreeNode) int {
	sum := 0

	var dfs5 func(root *TreeNode)
	dfs5 = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil && root.Left.Right == nil && root.Left.Left == nil {
			sum += root.Left.Val
		}
		dfs5(root.Left)
		dfs5(root.Right)
	}
	dfs5(root)
	return sum
}
