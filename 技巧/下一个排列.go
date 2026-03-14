func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 第一步，找到倒数第一个升序对
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		// 第二步，找到i后面比nums[i]大的最后一个元素
		j := n - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		// 交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 第三步，i后面区间逆序
	left, right := i+1, n-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}