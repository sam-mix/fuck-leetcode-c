package lib

func maxProfit2(prices []int) int {
	minPrice := 100000000
	profit := 0
	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		} else if profit < v-minPrice {
			profit = v - minPrice
		}
	}
	return profit
}

func maxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		v := prices[i]
		if v < minPrice {
			minPrice = v
		} else if profit < v-minPrice {
			profit = v - minPrice
		}
	}
	return profit
}
