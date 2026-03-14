package main

import "strings"

func decodeString(s string) string {
	// 保存进入 '[' 之前的字符串
	strStack := make([]string, 0)
	// 保存对应的重复次数 k
	numStack := make([]int, 0)

	cur := "" // 当前正在构建的字符串
	k := 0    // 当前累计的倍数（可能是多位数）

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch {
		case c >= '0' && c <= '9':
			// 处理多位数：例如 12[ab]
			k = k*10 + int(c-'0')

		case c == '[':
			// 进入括号：把当前状态压栈，然后重置
			numStack = append(numStack, k)
			strStack = append(strStack, cur)
			k = 0
			cur = ""

		case c == ']':
			// 出括号：弹出 k 和前缀字符串，把 cur 重复 k 次并拼接
			repeat := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			var b strings.Builder
			// 先写前缀，再写重复段
			b.WriteString(prev)
			for j := 0; j < repeat; j++ {
				b.WriteString(cur)
			}
			cur = b.String()

		default:
			// 普通字母：直接追加
			cur += string(c)
		}
	}

	return cur
}
