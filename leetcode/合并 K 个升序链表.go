package leetcode

func mergeKLists(lists []*ListNode) *ListNode {
	var ans *ListNode = nil
	for i := 0; i < len(lists); i++ {
		ans = mergeTwoList(ans, lists[i])
	}
	return ans
}

func mergeTwoList(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var res *ListNode
	// 如果res为空，res.Next就是空指针
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
