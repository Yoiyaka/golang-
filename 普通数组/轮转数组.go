func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // 防止k大于n
	// 1. 翻转整个数组
	reverse(nums, 0, n-1)
	// 2. 翻转前k个
	reverse(nums, 0, k-1)
	// 3. 翻转剩下的
	reverse(nums, k, n-1)
}

// 翻转nums[left:right]区间
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}