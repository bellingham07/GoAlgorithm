package tree

func levelOrder(root *TreeNode) [][]int {
	arr := [][]int{}
	depth := 0
	var order func(root *TreeNode, depth int)
	order = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], root.Val)
		order(root.Left, depth+1)
		order(root.Right, depth+1)
	}
	order(root, depth)
	// 自底向上输出结果 反转数组就行
	//for i := 0; i < len(arr)/2; i++ {
	//	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	//}
	return arr
}

func levelOrder2(root *TreeNode) [][]int {
	arr := [][]int{}
	var order func(node *TreeNode, depth int)
	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(arr) == depth {
			arr = append(arr, []int{})
		}
		arr[depth] = append(arr[depth], node.Val)
		order(node.Left, depth+1)
		order(node.Right, depth+1)
	}
	order(root, 0)
	return arr
}
