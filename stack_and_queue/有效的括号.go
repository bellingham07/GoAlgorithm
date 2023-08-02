package stack_and_queue

func isValid(s string) bool {
	hash := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			// 入栈
			stack = append(stack, s[i])
			// 栈内有元素并且栈顶元素和当前括号成对
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func ok(s string) bool {
	m := map[byte]byte{')': '(', ']': '[', '}': '{'}
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
