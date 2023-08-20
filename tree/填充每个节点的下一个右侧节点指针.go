package tree

type Node1 struct {
	Val   int
	Left  *Node1
	Right *Node1
	Next  *Node1
}

func connect(root *Node1) *Node1 {
	if root == nil {
		return nil
	}
	queue := []*Node1{root}
	for len(queue) > 0 {
		temp := queue
		queue = nil
		for i, node1 := range temp {
			if i+1 < len(temp) {
				node1.Next = temp[i+1]
			}
			// 拓展下一层节点
			if node1.Left != nil {
				queue = append(queue, node1.Left)
			}
			if node1.Right != nil {
				queue = append(queue, node1.Right)
			}
		}
	}
	return root
}
