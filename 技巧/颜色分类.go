func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	i := 0
	for i <= right {
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
			// 注意此处 i 不加1，因为交换来的是没检查过的数
		} else {
			i++
		}
	}
}