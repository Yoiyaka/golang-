package p0055_jump_game

func canJump(nums []int) bool {
	maxReach := 0
	for i, n := range nums {
		if i > maxReach {
			return false
		}
		if i+n > maxReach {
			maxReach = i + n
		}
	}
	return true
}
