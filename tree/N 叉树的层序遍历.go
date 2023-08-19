package tree

type Node struct {
	Val      int
	Children []*Node
}

// bfs
func levelOrder3(root *Node) [][]int {
	res := [][]int{} //结果集
	if root == nil {
		return res
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		t := []int{}
		for i := 0; i < l; i++ {
			node := queue[i]
			t = append(t, node.Val)
			if node.Children != nil {
				for j := 0; j < len(node.Children); j++ {
					queue = append(queue, node.Children[j])
				}
			}
		}
		res = append(res, t)
		queue = queue[l:]
	}

	return res
}

// dfs
func levelOrder4(root *Node) [][]int {
	res := [][]int{} //结果集
	if root == nil {
		return res
	}

	var dfs2 func(root *Node, deep int)
	dfs2 = func(root *Node, deep int) {
		if root == nil {
			return
		}
		// 第一次抵达这一层，添加元素
		if deep == len(res) {
			res = append(res, []int{})
		}
		res[deep] = append(res[deep], root.Val)
		if root.Children != nil {
			for i := 0; i < len(root.Children); i++ {
				dfs2(root.Children[i], deep+1)
			}
		}
	}
	dfs2(root, 0)
	return res
}
