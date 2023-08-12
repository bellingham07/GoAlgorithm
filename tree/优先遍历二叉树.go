package tree

import (
	"container/list"
	"fmt"
)

func NewNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

func DFS(root *TreeNode) {
	if root == nil {
		return
	}

	DFS(root.Left)

	DFS(root.Right)
	fmt.Printf("%d ", root.Val)

}

func BFS(root *TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	// 将头节点放到队尾
	queue.PushBack(root)

	for queue.Len() > 0 {
		// 从队首取到元素
		node := queue.Remove(queue.Front()).(*TreeNode)
		fmt.Printf("%d ", node.Val)

		// 如果子节点不为空，将其放入队尾
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}

func main() {
	/*
	       1
	      / \
	     2   3
	    / \
	   4   5
	*/
	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)

	fmt.Println("DFS:")
	DFS(root)
	fmt.Println("\nBFS:")
	BFS(root)
}
