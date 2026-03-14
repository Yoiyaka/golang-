func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
// rob 函数用于计算不触动警报装置情况下，一夜之内能够偷窃到的最高金额。
// 给定一个代表每个房屋存放金额的非负整数数组 nums。
func rob(nums []int) int {
    // prev1 保存到上一个房屋能获得的最高金额
    // prev2 保存到上上个房屋能获得的最高金额
    prev1, prev2 := 0, 0

    // 遍历每个房屋的金额
    for _, num := range nums {
        // 当前的最大金额等于：
        // 1. 不偷当前房屋，最大金额是 prev1；
        // 2. 偷当前房屋，最大金额是 prev2 + 当前房屋金额 num；
        // 取两者中的最大值
        curr := max(prev1, prev2+num)
        // 更新 prev2 和 prev1，为下一个循环做准备
        prev2 = prev1    // 上上家房变成上家房
        prev1 = curr     // 当前最大金额变成上家房
    }
    // 返回最后的最大金额
    return prev1
}
*/