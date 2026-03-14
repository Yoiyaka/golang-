package main

func isValid(s string) bool {
	// 右括号 -> 期望匹配的左括号
	match := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 用 slice 当栈
	stack := make([]byte, 0, len(s))

	for i := 0; i < len(s); i++ {
		c := s[i]

		// 左括号入栈
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			continue
		}

		// 右括号：栈不能为空，且栈顶必须匹配
		left, ok := match[c]
		if !ok { // 题目保证只有括号字符，这里算防御式
			return false
		}
		if len(stack) == 0 || stack[len(stack)-1] != left {
			return false
		}

		// 弹栈
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
