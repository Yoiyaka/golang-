package p0309_best_time_to_buy_and_sell_stock_with_cooldown

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	hold, sold, rest := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		prevHold, prevSold, prevRest := hold, sold, rest
		hold = max(prevHold, prevRest-prices[i])
		sold = prevHold + prices[i]
		rest = max(prevRest, prevSold)
	}
	if sold > rest {
		return sold
	}
	return rest
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
