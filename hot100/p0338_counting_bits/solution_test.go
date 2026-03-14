package p0338_counting_bits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tc := range tests {
		got := countBits(tc.n)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("countBits(%d) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
