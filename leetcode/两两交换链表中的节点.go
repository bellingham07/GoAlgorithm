package leetcode

func swapPairs(head *ListNode) *ListNode {
	// 虚拟头节点
	dummyHead := &ListNode{
		Val:  0,
		Next: head,
	}

	// 临时节点
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		// 交换节点
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		// temp指向最新的节点
		temp = node1
	}
	return dummyHead.Next
}
