package p0128_longest_consecutive_sequence

import "testing"

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{[]int{}, 0},
	}
	for _, tc := range tests {
		got := longestConsecutive(tc.nums)
		if got != tc.want {
			t.Errorf("longestConsecutive(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
