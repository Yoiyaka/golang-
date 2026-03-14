package p0169_majority_element

import "testing"

func TestMajorityElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tc := range tests {
		if got := majorityElement(tc.nums); got != tc.want {
			t.Errorf("majorityElement(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
