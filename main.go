package main

func main() {
	root := new(Node)
	root.Val = 1
	root.Left = new(Node)
	root.Right = new(Node)
	root.Left.Val = 2
	root.Right.Val = 3
	root.Left.Left.Val = 4
	root.Left.Right.Val = 5
	root.Right.Left.Val = 6
	root.Right.Right.Val = 7

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		l := len(queue)
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		for i := 1; i < l; i++ {
			node := queue[i-1]
			node1 := queue[i]
			node.Next = node1
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
