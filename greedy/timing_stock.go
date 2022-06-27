package greedy

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
func maxProfit(prices []int) int {
	var profit int
	var minPrice = prices[0]

	for _, currentPrice := range prices {
		if currentPrice < minPrice {
			minPrice = currentPrice
		} else {
			profit = max(profit, currentPrice-minPrice)
		}
	}

	return profit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
