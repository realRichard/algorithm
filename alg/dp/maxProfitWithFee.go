package dp

// 力扣第 714 题
// 买卖股票的最佳时机含手续费

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/description/

// 每次交易要支付手续费，只要把手续费从利润中减去即可

func MaxProfitWithFee(prices []int, fee int) int {
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
			// base case
			table[i][0] = 0
			table[i][1] = -prices[i] - fee
			continue
		}
		// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
		// 解释：相当于买入股票的价格升高了。
		// 在第一个式子里减也是一样的，相当于卖出股票的价格减小了。
		if table[i-1][0] > table[i-1][1]+prices[i] {
			table[i][0] = table[i-1][0]
		} else {
			table[i][0] = table[i-1][1] + prices[i]
		}
		if table[i-1][1] > table[i-1][0]-prices[i]-fee {
			table[i][1] = table[i-1][1]
		} else {
			table[i][1] = table[i-1][0] - prices[i] - fee
		}
	}
	return table[n-1][0]
}
