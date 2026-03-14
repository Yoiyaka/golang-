package p0312_burst_balloons

func maxCoins(nums []int) int {
	n := len(nums)
	balloons := make([]int, n+2)
	balloons[0] = 1
	balloons[n+1] = 1
	for i := 0; i < n; i++ {
		balloons[i+1] = nums[i]
	}
	n += 2
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for length := 2; length < n; length++ {
		for left := 0; left < n-length; left++ {
			right := left + length
			for k := left + 1; k < right; k++ {
				coins := balloons[left]*balloons[k]*balloons[right] + dp[left][k] + dp[k][right]
				if coins > dp[left][right] {
					dp[left][right] = coins
				}
			}
		}
	}
	return dp[0][n-1]
}
