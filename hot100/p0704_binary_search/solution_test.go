package p0704_binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{5}, 5, 0},
	}
	for _, tc := range tests {
		got := search(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("search(%v,%d) = %d, want %d", tc.nums, tc.target, got, tc.want)
		}
	}
}
