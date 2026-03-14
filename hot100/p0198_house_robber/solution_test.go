package p0198_house_robber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for _, tc := range tests {
		if got := rob(tc.nums); got != tc.want {
			t.Errorf("rob(%v) = %d, want %d", tc.nums, got, tc.want)
		}
	}
}
