package p0239_sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
	}
	for _, tc := range tests {
		got := maxSlidingWindow(tc.nums, tc.k)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("maxSlidingWindow(%v,%d) = %v, want %v", tc.nums, tc.k, got, tc.want)
		}
	}
}
