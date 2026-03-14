package p0045_jump_game_ii

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
	}
	for _, tc := range tests {
		if got := jump(tc.nums); got != tc.want {
			t.Errorf("jump(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
