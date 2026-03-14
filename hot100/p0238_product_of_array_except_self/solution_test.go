package p0238_product_of_array_except_self

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tc := range tests {
		got := productExceptSelf(tc.nums)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("productExceptSelf(%v) = %v, want %v", tc.nums, got, tc.want)
		}
	}
}
