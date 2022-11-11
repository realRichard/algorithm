package dp

// 力扣第 309 题
// 最佳买卖股票时机含冷冻期

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/

func MaxProfitWithCold(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	table := make([][]int, len(prices))
	for index := range table {
		table[index] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			// base case 1
			table[i][0] = 0
			table[i][1] = -prices[0]
			continue
		}
		if i-2 == -1 {
			// base case 2
			if table[i-1][0] > table[i-1][1]+prices[i] {
				table[i][0] = table[i-1][0]
			} else {
				table[i][0] = table[i-1][1] + prices[i]
			}
			if table[i-1][1] > table[i-1][0]-prices[i] {
				table[i][1] = table[i-1][1]
			} else {
				table[i][1] = table[i-1][0] - prices[i]
			}
			continue
		}
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		// dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
		// 解释：第 i 天选择 buy 的时候，要从 i-2 的状态转移，而不是 i-1 。
		if table[i-1][0] > table[i-1][1]+prices[i] {
			table[i][0] = table[i-1][0]
		} else {
			table[i][0] = table[i-1][1] + prices[i]
		}
		if table[i-1][1] > table[i-2][0]-prices[i] {
			table[i][1] = table[i-1][1]
		} else {
			table[i][1] = table[i-2][0] - prices[i]
		}
	}
	return table[n-1][0]
}
