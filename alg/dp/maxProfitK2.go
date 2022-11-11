package dp

// 力扣第 123 题
// 买卖股票的最佳时机 III

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/

func MaxProfitK2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	maxK := 2
	table := make([][][]int, n)
	for index := range table {
		table[index] = make([][]int, maxK+1)
	}
	for i := range table {
		for j := range table[i] {
			table[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for k := maxK; k > 0; k-- {
			if i-1 == -1 {
				// base case
				table[i][k][0] = 0
				table[i][k][1] = -prices[i]
				continue
			}
			if table[i-1][k][0] > table[i-1][k][1]+prices[i] {
				table[i][k][0] = table[i-1][k][0]
			} else {
				table[i][k][0] = table[i-1][k][1] + prices[i]
			}
			if table[i-1][k][1] > table[i-1][k-1][0]-prices[i] {
				table[i][k][1] = table[i-1][k][1]
			} else {
				table[i][k][1] = table[i-1][k-1][0] - prices[i]
			}
		}
	}
	return table[n-1][maxK][0]
}
