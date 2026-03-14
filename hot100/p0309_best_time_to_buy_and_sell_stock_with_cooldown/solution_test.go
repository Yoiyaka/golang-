package p0309_best_time_to_buy_and_sell_stock_with_cooldown

import "testing"

func TestMaxProfit(t *testing.T) {
	if got := maxProfit([]int{1, 2, 3, 0, 2}); got != 3 {
		t.Errorf("got %d, want 3", got)
	}
	if got := maxProfit([]int{1}); got != 0 {
		t.Errorf("got %d, want 0", got)
	}
}
