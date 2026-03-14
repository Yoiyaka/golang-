func coinChange(coins []int, amount int) int {
	// dp[i] 表示凑成金额 i 的最少硬币数
	dp := make([]int, amount+1)
	// 初始化：最大值（假设币数不会超额），用 amount+1 作为无解值
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0 // 金额为 0 时，不需要任何硬币

	// 自底向上的动态规划
	for i := 1; i <= amount; i++ {
		// 遍历每种硬币面额
		for _, coin := range coins {
			if i >= coin {
				// 如果当前金额大于等于该硬币面额，则可以转移状态
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	// 如果无法凑成，总币数为 amount+1，返回 -1
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}