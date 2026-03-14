package p0034_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}
	for _, tc := range tests {
		got := searchRange(tc.nums, tc.target)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("searchRange(%v,%d) = %v, want %v", tc.nums, tc.target, got, tc.want)
		}
	}
}
