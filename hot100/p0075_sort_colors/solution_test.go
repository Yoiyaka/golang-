package p0075_sort_colors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	want := []int{0, 0, 1, 1, 2, 2}
	if !reflect.DeepEqual(nums, want) {
		t.Errorf("got %v, want %v", nums, want)
	}
	nums2 := []int{2, 0, 1}
	sortColors(nums2)
	want2 := []int{0, 1, 2}
	if !reflect.DeepEqual(nums2, want2) {
		t.Errorf("got %v, want %v", nums2, want2)
	}
}
