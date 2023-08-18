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
	//deep==len(*res)时候。说明第一次访问到这一层，将元素加入
	if deep == len(*res) {
		*res = append(*res, node.Val)
	}
	// 要确保该节点子节点不为空，才能进行下一次递归，否则访问了一个空节点，也是第一次到达那一层，此时，node为空，使用node.Val会出现空指针
	if node.Right != nil {
		dfs(node.Right, deep+1, res)
	}
	if node.Left != nil {
		dfs(node.Left, deep+1, res)
	}
}

// bfs
func rightSideView2(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		deep := len(queue)
		var rightNode *TreeNode
		for i := 0; i < deep; i++ {
			node := queue[i]
			rightNode = node
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[deep:]
		res = append(res, rightNode.Val)
	}
	return res
}
