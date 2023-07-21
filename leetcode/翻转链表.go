package leetcode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	pre = nil
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	head = cur
	return pre
}
