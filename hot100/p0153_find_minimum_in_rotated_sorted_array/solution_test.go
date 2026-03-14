package p0153_find_minimum_in_rotated_sorted_array

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}
	for _, tc := range tests {
		got := findMin(tc.nums)
		if got != tc.want {
			t.Errorf("findMin(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
