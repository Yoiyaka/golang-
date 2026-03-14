package p0300_longest_increasing_subsequence

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7}, 1},
	}
	for _, tc := range tests {
		if got := lengthOfLIS(tc.nums); got != tc.want {
			t.Errorf("lengthOfLIS(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
