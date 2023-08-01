package stack_and_queue

//将一个栈当作输入栈，用于压入 push\texttt{push}push 传入的数据；另一个栈当作输出栈，用于 pop\texttt{pop}pop 和 peek\texttt{peek}peek 操作。
//
//每次 pop\texttt{pop}pop 或 peek\texttt{peek}peek 时，若输出栈为空则将输入栈的全部数据依次弹出并压入输出栈，这样输出栈从栈顶往栈底的顺序就是队列从队首往队尾的顺序

type MyQueue struct {
	inStack, outStack []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (q *MyQueue) Push(x int) {
	// 入栈操作，直接使用append将元素放到队列末尾
	q.inStack = append(q.inStack, x)
}

func (q *MyQueue) in2out() {
	// 将输入栈元素转移到输出栈
	for len(q.inStack) > 0 {
		q.outStack = append(q.outStack, q.inStack[len(q.inStack)-1])
		q.inStack = q.inStack[:len(q.inStack)-1]
	}
}

func (q *MyQueue) Pop() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	x := q.outStack[len(q.outStack)-1]
	q.outStack = q.outStack[:len(q.outStack)-1]
	return x
}

func (q *MyQueue) Peek() int {
	if len(q.outStack) == 0 {
		q.in2out()
	}
	return q.outStack[len(q.outStack)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.inStack) == 0 && len(q.outStack) == 0
}
