package tree

// bfs
func largestValues(root *TreeNode) (ans []int) {
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		temp := queue[0].Val
		for i := 0; i < l; i++ {
			node := queue[i]
			if node.Val > temp {
				temp = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, temp)
		queue = queue[l:]
	}
	return ans
}

// dfs
func largestValues2(root *TreeNode) (ans []int) {
	var dfs5 func(*TreeNode, int)
	dfs5 = func(node *TreeNode, curHeight int) {
		if node == nil {
			return
		}
		if curHeight == len(ans) {
			ans = append(ans, node.Val)
		} else {
			ans[curHeight] = max(ans[curHeight], node.Val)
		}
		dfs5(node.Left, curHeight+1)
		dfs5(node.Right, curHeight+1)
	}
	dfs5(root, 0)
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
