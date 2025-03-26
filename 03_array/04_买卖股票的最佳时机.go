package array

// MaxProfitOnce 买卖股票的最佳时机，只允许购买一次
func MaxProfitOnce(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ret := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		ret = max(ret, prices[i]-minPrice)
	}
	return ret
}

// MaxProfitNoLimit 买卖股票的最佳时机，不限制购买次数
// 每次上涨购买赎回即可
func MaxProfitNoLimit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ret := 0
	for i := 1; i < len(prices); i++ {
		ret += max(0, prices[i]-prices[i-1])
	}
	return ret
}
