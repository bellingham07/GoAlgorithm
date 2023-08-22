package list

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var res *ListNode
	// res为空，会出现空指针
	res = new(ListNode)
	head := res

	for {
		// 这里如果判断list1.Next==nil 会导致无法检查下一个元素
		if list1 == nil {
			res.Next = list2
			res = res.Next
			break
		}
		if list2 == nil {
			res.Next = list1
			res = res.Next
			break
		}
		if list1.Val > list2.Val {
			res.Next = list2
			res = res.Next
			list2 = list2.Next
		} else {
			res.Next = list1
			res = res.Next
			list1 = list1.Next
		}
	}
	return head.Next
}
