func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	res := 0
	if len(prices) == 0 {
		return 0
	}
	var lowestPriceSoFar int = prices[0]
	for i := 1; i < len(prices); i++ {
		var maxProfitSellToday int = prices[i] - lowestPriceSoFar
		res = max(res, maxProfitSellToday)
		lowestPriceSoFar = min(lowestPriceSoFar, prices[i])
	}
	return res
}