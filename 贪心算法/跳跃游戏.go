package main

func canJump(nums []int) bool {
	far := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i > far {
			return false
		}
		if i+nums[i] > far {
			far = i + nums[i]
		}
		if far >= n-1 {
			return true
		}
	}
	return true
}
