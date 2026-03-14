func generateParenthesis(n int) []string {
	var res []string
	var dfs func(cur string, left int, right int)
	dfs = func(cur string, left int, right int) {
		if left == 0 && right == 0 {
			res = append(res, cur)
			return
		}
		if left > 0 {
			dfs(cur+"(", left-1, right)
		}
		if right > 0 && right > left {
			dfs(cur+")", left, right-1)
		}
	}
	dfs("", n, n)
	return res
}