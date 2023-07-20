package leetcode

type MyLinkedList struct {
	Size int
	Next *ListNode
}

func constructor() MyLinkedList {
	return MyLinkedList{0, &ListNode{}}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.Size {
		return -1
	}
	cur := this.Next
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.Size, val)

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}
	index = max(index, 0)
	this.Size++
	pred := this.Next
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	toAdd := &ListNode{val, pred.Next}
	pred.Next = toAdd
}
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	this.Size--
	pred := this.Next
	for i := 0; i < index; i++ {
		pred = pred.Next
	}
	pred.Next = pred.Next.Next
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
