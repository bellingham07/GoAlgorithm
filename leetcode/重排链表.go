package leetcode

// 找到链表中点，断开链表，反转第二个链表，分别合并两个链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	middle := mid(head)

	l1 := head
	l2 := middle.Next
	middle.Next = nil

	l2 = reverse(l2)

	merge(l1, l2)
}

func mid(l *ListNode) *ListNode {
	// 快慢指针，快指针走完，慢指针正好在中间位置
	slow, fast := l, l.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func reverse(l *ListNode) *ListNode {
	// 反转链表，三个指针，prev指向前一个地址，cur当前位置，nextTmp下一个位置
	var prev, cur *ListNode = nil, l
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	// 返回prev是因为，cur==nil的时候，prev正好在队首的位置，也就是头指针
	return prev
}

func merge(l1, l2 *ListNode) {
	// 交替合并链表
	for l1 != nil && l2 != nil {
		// 先记录两个指针的下一个节点
		l1temp := l1.Next
		l2temp := l2.Next

		// 断开指针，next指向新地址
		l1.Next = l2
		l1 = l1temp

		l2.Next = l1
		l2 = l2temp
	}
}
