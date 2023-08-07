package stack_and_queue

import "container/list"

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

func isValid2(s string) bool {
	l := list.New()
	valid := map[any]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	l.PushFront('?')
	for _, str := range s {
		l.PushBack(str)
		if str == ']' || str == '}' || str == ')' {
			if valid[l.Back().Prev().Value] != l.Back().Value {
				return false
			} else {
				l.Remove(l.Back())
				l.Remove(l.Back())
			}
		}
	}
	if l.Len() > 1 {
		return false
	}
	return true
}
