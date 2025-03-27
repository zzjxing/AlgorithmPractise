package array

import "math"

// CoinChange 零钱兑换，经典 dp 问题
// dp[i]: 兑换 i 所需要的最小硬币数量
// 状态转移方程，dp[i]=min(dp[i],dp[i-coin]+1)
func CoinChange(coins []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[target] == math.MaxInt32 {
		return -1
	}
	return dp[target]
}
