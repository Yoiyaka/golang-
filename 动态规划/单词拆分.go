import "strings"

// wordBreak 函数用于判断字符串 s 能否由 wordDict 中的单词拼接而成。
func wordBreak(s string, wordDict []string) bool {
	// 将 wordDict 转为 map，以便 O(1) 查询
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// dp[i] 表示 s[0:i] 是否可以由字典单词拼接
	dp := make([]bool, len(s)+1)
	dp[0] = true // 空字符串可以拼接

	// 遍历每个 s 的前缀
	for i := 1; i <= len(s); i++ {
		// 检查 s 中 [j:i) 是否在词典，且 s[0:j] 可被拆分
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	// 返回是否能拼接出 s
	return dp[len(s)]
}