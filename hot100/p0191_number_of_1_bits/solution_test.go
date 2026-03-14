package p0191_number_of_1_bits

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{11, 3},           // 00000000000000000000000000001011
		{128, 1},          // 00000000000000000000000010000000
		{4294967293, 31},  // 11111111111111111111111111111101
	}
	for _, tc := range tests {
		if got := hammingWeight(tc.num); got != tc.want {
			t.Errorf("hammingWeight(%d) = %d, want %d", tc.num, got, tc.want)
		}
	}
}
