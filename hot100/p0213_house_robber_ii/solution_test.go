package p0213_house_robber_ii

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 2}, 3},
		{[]int{1, 2, 3, 1}, 4},
		{[]int{1, 2, 3}, 3},
	}
	for _, tc := range tests {
		if got := rob(tc.nums); got != tc.want {
			t.Errorf("rob(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
