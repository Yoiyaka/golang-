package p0053_maximum_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}
	for _, tc := range tests {
		if got := maxSubArray(tc.nums); got != tc.want {
			t.Errorf("maxSubArray(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
