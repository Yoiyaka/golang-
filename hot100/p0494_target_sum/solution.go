package p0494_target_sum

func findTargetSumWays(nums []int, target int) int {
	memo := make(map[[2]int]int)
	var dp func(i, currSum int) int
	dp = func(i, currSum int) int {
		if i == len(nums) {
			if currSum == target {
				return 1
			}
			return 0
		}
		key := [2]int{i, currSum}
		if v, ok := memo[key]; ok {
			return v
		}
		result := dp(i+1, currSum+nums[i]) + dp(i+1, currSum-nums[i])
		memo[key] = result
		return result
	}
	return dp(0, 0)
}
