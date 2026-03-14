package p0121_best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}
