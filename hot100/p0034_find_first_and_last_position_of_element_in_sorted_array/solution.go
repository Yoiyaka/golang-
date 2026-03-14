package p0034_find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	first := findFirst(nums, target)
	last := findLast(nums, target)
	return []int{first, last}
}

func findFirst(nums []int, target int) int {
	left, right, result := 0, len(nums)-1, -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func findLast(nums []int, target int) int {
	left, right, result := 0, len(nums)-1, -1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
