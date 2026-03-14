package p0322_coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 5, 10, 25}, 30, 2},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}
	for _, tc := range tests {
		if got := coinChange(tc.coins, tc.amount); got != tc.want {
			t.Errorf("coinChange(%v,%d) = %d, want %d", tc.coins, tc.amount, got, tc.want)
		}
	}
}
