package p0283_move_zeroes

func moveZeroes(nums []int) {
	pos := 0
	for _, v := range nums {
		if v != 0 {
			nums[pos] = v
			pos++
		}
	}
	for pos < len(nums) {
		nums[pos] = 0
		pos++
	}
}
