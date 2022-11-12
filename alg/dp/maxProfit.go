package dp

// 力扣第 121 题
// 买卖股票的最佳时机

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/

// https://labuladong.github.io/algo/1/13/#三秒杀题目
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems

func MaxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	table := make([][]int, n)
	for index := range table {
		table[index] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i-1 < 0 {
			table[0][0] = 0
			table[0][1] = -prices[0]
		} else {
			if table[i-1][0] > table[i-1][1]+prices[i] {
				table[i][0] = table[i-1][0]
			} else {
				table[i][0] = table[i-1][1] + prices[i]
			}
			if table[i-1][1] > -prices[i] {
				table[i][1] = table[i-1][1]
			} else {
				table[i][1] = -prices[i]
			}
		}
	}
	return table[n-1][0]
}
