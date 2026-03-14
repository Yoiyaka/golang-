package p0215_kth_largest_element_in_an_array

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tc := range tests {
		got := findKthLargest(tc.nums, tc.k)
		if got != tc.want {
			t.Errorf("findKthLargest(%v,%d) = %d, want %d", tc.nums, tc.k, got, tc.want)
		}
	}
}
