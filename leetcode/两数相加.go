package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 进位
	carry := 0

	head := &ListNode{}
	p := head

	for l1 != nil || l2 != nil {
		temp := carry
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += l2.Val
			l2 = l2.Next
		}
		p.Next = &ListNode{Val: temp % 10}
		carry = temp / 10
		p = p.Next
	}
	if carry != 0 {
		p.Next = &ListNode{Val: carry}
	}
	return head.Next
}
