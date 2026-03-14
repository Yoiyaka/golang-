package p0084_largest_rectangle_in_histogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{[]int{1}, 1},
		{[]int{1, 1}, 2},
	}
	for _, tc := range tests {
		got := largestRectangleArea(tc.heights)
		if got != tc.want {
			t.Errorf("largestRectangleArea(%v) = %d, want %d", tc.heights, got, tc.want)
		}
	}
}
