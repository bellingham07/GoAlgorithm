package tree

type data struct {
	sum, count int
}

// dfs
func averageOfLevels(root *TreeNode) []float64 {
	res := []data{}
	var dfs2 func(node *TreeNode, depth int)
	dfs2 = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		// 什么情况下depth会小于len（res）？回溯的时候
		if depth < len(res) {
			res[depth].sum += node.Val
			res[depth].count++
		} else {
			// 第一次到达这一层的时候，需要手动添加节点
			res = append(res, data{
				sum:   node.Val,
				count: 1,
			})
		}
		dfs2(node.Left, depth+1)
		dfs2(node.Right, depth+1)
	}
	dfs2(root, 0)

	ans := make([]float64, 0)

	for i, d := range res {
		ans[i] = float64(d.sum) / float64(d.count)
	}

	return ans
}

//bfs

func averageOfLevels2(root *TreeNode) []float64 {
	ans := []float64{}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		sum := 0
		for i := 0; i < l; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		num := float64(sum) / float64(l)
		ans = append(ans, num)
		queue = queue[l:]
	}

	return ans
}
