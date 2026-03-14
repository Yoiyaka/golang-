package p0031_next_permutation

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}
	for _, tc := range tests {
		nextPermutation(tc.input)
		if !reflect.DeepEqual(tc.input, tc.want) {
			t.Errorf("got %v, want %v", tc.input, tc.want)
		}
	}
}
