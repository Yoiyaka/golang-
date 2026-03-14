package p0213_house_robber_ii

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, len(nums)-2), robRange(nums, 1, len(nums)-1))
}

func robRange(nums []int, start, end int) int {
	prev2, prev1 := 0, 0
	for i := start; i <= end; i++ {
		curr := prev1
		if prev2+nums[i] > curr {
			curr = prev2 + nums[i]
		}
		prev2, prev1 = prev1, curr
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
