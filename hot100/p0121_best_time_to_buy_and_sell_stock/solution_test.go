package p0121_best_time_to_buy_and_sell_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, tc := range tests {
		if got := maxProfit(tc.prices); got != tc.want {
			t.Errorf("maxProfit(%v) = %d, want %d", tc.prices, got, tc.want)
		}
	}
}
