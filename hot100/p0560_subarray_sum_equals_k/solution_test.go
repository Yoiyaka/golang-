package p0560_subarray_sum_equals_k

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1}, 1, 1},
	}
	for _, tc := range tests {
		got := subarraySum(tc.nums, tc.k)
		if got != tc.want {
			t.Errorf("subarraySum(%v,%d) = %d, want %d", tc.nums, tc.k, got, tc.want)
		}
	}
}
