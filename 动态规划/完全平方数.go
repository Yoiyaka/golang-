func numSquares(n int) int {
	// dp[i] 表示数字 i 最少可以由多少个完全平方数组成
	dp := make([]int, n+1)
	dp[0] = 0 // 0 不需要任何平方数

	// 初始化，每个数最坏都是用1组成，比如 7=1+1+...+1
	for i := 1; i <= n; i++ {
		dp[i] = i // 最多使用 i 个 1
		// 遍历所有小于等于i的完全平方数
		for j := 1; j*j <= i; j++ {
			// 用 j*j，剩下部分是 dp[i - j*j]
			if dp[i-j*j]+1 < dp[i] {
				dp[i] = dp[i-j*j] + 1
			}
		}
	}
	return dp[n]
}