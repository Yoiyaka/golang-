func partition(s string) [][]string {
	var res [][]string
	var path []string

	// 判断s[l:r+1]是不是回文串
	isPalindrome := func(s string, left, right int) bool {
		for left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	var dfs func(int)
	dfs = func(start int) {
		if start == len(s) {
			// 深拷贝path
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for end := start; end < len(s); end++ {
			if isPalindrome(s, start, end) {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return res
}