package offer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList1(head *Node) *Node {
	// 复制next
	temp := head

	var ans *Node
	ans = new(Node)
	temp2 := ans
	for temp != nil {
		node := &Node{Val: temp.Val}
		temp2.Next = node
		temp2 = temp2.Next
		temp = temp.Next
	}

	temp = head
	temp2 = ans.Next

	for temp != nil {
		v := temp.Random
		if v == nil {
			temp2.Random = nil
		} else {
			t := ans.Next
			for t.Next != nil {
				if t.Val == v.Val && t.Next.Val == v.Next.Val {
					temp2.Random = t
					break
				}
				t = t.Next
			}
		}
		temp = temp.Next
		temp2 = temp2.Next
	}
	return ans.Next
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 第一步：复制每个节点，并将复制的节点插入原节点之后
	current := head
	for current != nil {
		copyNode := &Node{Val: current.Val}
		copyNode.Next = current.Next
		current.Next = copyNode
		current = copyNode.Next
	}

	// 第二步：设置复制节点的 Random 指针
	current = head
	for current != nil {
		if current.Random != nil {
			current.Next.Random = current.Random.Next
		}
		current = current.Next.Next
	}

	// 第三步：拆分原链表和复制链表
	current = head
	newHead := head.Next
	for current != nil {
		copyNode := current.Next
		current.Next = copyNode.Next
		if copyNode.Next != nil {
			copyNode.Next = copyNode.Next.Next
		}
		current = current.Next
	}

	return newHead
}
