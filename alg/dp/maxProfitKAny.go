package dp

// 力扣第 188 题
// 买卖股票的最佳时机 IV

func MaxProfitKAny(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	// 一次交易由买入和卖出构成，至少需要两天。
	// 所以说有效的限制 k 应该不超过 n/2，如果超过，就没有约束作用了，相当于 k 没有限制的情况，而这种情况是之前解决过的
	if k > n/2 {
		return MaxProfitKInfV2(prices)
	}
	table := make([][][]int, n)
	for index := range table {
		table[index] = make([][]int, k+1)
	}
	for i := range table {
		for j := range table[i] {
			table[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for j := k; j > 0; j-- {
			if i-1 == -1 {
				// base case
				table[i][j][0] = 0
				table[i][j][1] = -prices[i]
				continue
			}
			if table[i-1][j][0] > table[i-1][j][1]+prices[i] {
				table[i][j][0] = table[i-1][j][0]
			} else {
				table[i][j][0] = table[i-1][j][1] + prices[i]
			}
			if table[i-1][j][1] > table[i-1][j-1][0]-prices[i] {
				table[i][j][1] = table[i-1][j][1]
			} else {
				table[i][j][1] = table[i-1][j-1][0] - prices[i]
			}
		}
	}
	return table[n-1][k][0]
}
