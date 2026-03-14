func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxVal, minVal := nums[0], nums[0]
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		tempMax := maxVal
		maxVal = max(nums[i], max(nums[i]*maxVal, nums[i]*minVal))
		minVal = min(nums[i], min(nums[i]*tempMax, nums[i]*minVal))
		result = max(result, maxVal)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}