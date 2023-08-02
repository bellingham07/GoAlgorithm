package stack_and_queue

func removeDuplicates(s string) string {
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if i == 0 {
			stack = append(stack, s[i])
			// 只有栈不为空的情况下才能修改
		} else if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return string(stack)
}
