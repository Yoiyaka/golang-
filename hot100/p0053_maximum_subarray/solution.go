package p0053_maximum_subarray

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if curr < 0 {
			curr = nums[i]
		} else {
			curr += nums[i]
		}
		if curr > maxSum {
			maxSum = curr
		}
	}
	return maxSum
}
