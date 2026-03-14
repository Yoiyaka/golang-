package p0283_move_zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 0, 1}, []int{1, 1, 0}},
	}
	for _, tc := range tests {
		moveZeroes(tc.input)
		if !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("moveZeroes result = %v, want %v", tc.input, tc.want)
		}
	}
}
