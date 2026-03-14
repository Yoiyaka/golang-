package p0268_missing_number

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
	}
	for _, tc := range tests {
		if got := missingNumber(tc.nums); got != tc.want {
			t.Errorf("missingNumber(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
