package tree

// dfs
func maxDepth(root *TreeNode) int {
	res := []int{}
	var dfs func(root *TreeNode, deep int)
	dfs = func(root *TreeNode, deep int) {
		if root == nil {
			return
		}
		if deep == len(res) {
			res = append(res, 1)
		}
		deep++
		dfs(root.Left, deep)
		dfs(root.Right, deep)
	}
	dfs(root, 0)
	return len(res)
}

// bfs

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	ans := 0
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[l:]
		ans++
	}
	return ans
}
