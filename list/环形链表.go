package list

func hasCycle(head *ListNode) bool {
	// 确保初始节点不为空
	if head == nil || head.Next == nil {
		return false
	}
	// 用快慢指针
	fast, slow := head.Next, head
	// 两个指针相遇，说明有环
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}
