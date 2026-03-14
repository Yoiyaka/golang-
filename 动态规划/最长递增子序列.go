func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	res := 1
	// 初始化每个 dp[i] 为 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			// 如果 nums[i] 严格大于 nums[j]，则可以递增
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func lengthOfLIS(nums []int) int {
    tails := []int{}
    for _, num := range nums {
        left, right := 0, len(tails)
        for left < right {
            mid := (left + right) / 2
            if tails[mid] < num {
                left = mid + 1
            } else {
                right = mid
            }
        }
        if left == len(tails) {
            tails = append(tails, num)
        } else {
            tails[left] = num
        }
    }
    return len(tails)
}
*/