package p0215_kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	pivot := nums[right]
	p := left
	for i := left; i < right; i++ {
		if nums[i] <= pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	nums[p], nums[right] = nums[right], nums[p]
	if p == k {
		return nums[p]
	} else if p < k {
		return quickSelect(nums, p+1, right, k)
	}
	return quickSelect(nums, left, p-1, k)
}
