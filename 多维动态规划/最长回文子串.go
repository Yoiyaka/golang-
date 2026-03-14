func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	start, maxLen := 0, 1

	expand := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			if right-left+1 > maxLen {
				start = left
				maxLen = right - left + 1
			}
			left--
			right++
		}
	}

	for i := 0; i < len(s); i++ {
		expand(i, i)   // 奇数长度中心
		expand(i, i+1) // 偶数长度中心
	}
	return s[start : start+maxLen]
}

/*manacher算法
func longestPalindrome(s string) string {
    // 预处理
    t := "^"
    for _, c := range s {
        t += "#" + string(c)
    }
    t += "#$"

    n := len(t)
    p := make([]int, n)
    center, right := 0, 0
    maxLen, maxCenter := 0, 0

    for i := 1; i < n-1; i++ {
        mirror := 2*center - i
        if i < right {
            p[i] = min(right-i, p[mirror])
        }
        // 尽量向两边扩展
        for t[i+p[i]+1] == t[i-p[i]-1] {
            p[i]++
        }
        // 更新回文右边界和中心
        if i+p[i] > right {
            center = i
            right = i + p[i]
        }
        if p[i] > maxLen {
            maxLen = p[i]
            maxCenter = i
        }
    }
    start := (maxCenter - maxLen) / 2
    return s[start : start+maxLen]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
*/