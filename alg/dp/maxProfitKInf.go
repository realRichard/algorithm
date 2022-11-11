package dp

// 力扣第 122 题
// 买卖股票的最佳时机 II

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/

func MaxProfitKInfV1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 递增区间左右指针
	left, right := 0, 0
	// 因为可以交易多次
	// 根据股票价格折线图，每一次递增区间幅度相加就是最大利润了
	maxProfit := 0
	// 记录递增区间的差值
	diff := 0
	for right < len(prices)-1 {
		// 递增
		if prices[right] <= prices[right+1] {
			right++
			diff = prices[right] - prices[left]
		} else {
			// 递减
			maxProfit += diff
			diff = 0
			left, right = right+1, right+1
		}
	}
	// 最后一段递增区间也要加上，如果最后一段是递减区间则 diff = 0
	maxProfit += diff
	return maxProfit
}

func MaxProfitKInfV2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	// 没有交易总数的限制，也就相当于 k 为正无穷了
	// 如果 k 为正无穷，那么就可以认为 k 和 k - 1 是一样的

	// dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
	// dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])
	// 			   = max(dp[i-1][k][1], dp[i-1][k][0] - prices[i])

	// 我们发现数组中的 k 已经不会改变了，也就是说不需要记录 k 这个状态了：
	// dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
	// dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])

	table := make([][]int, len(prices))
	for index := range table {
		table[index] = make([]int, 2)
	}
	for i := 0; i < len(prices); i++ {
		if i-1 < 0 {
			// base case
			table[i][0] = 0
			table[i][1] = -prices[0]
			continue
		}
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
	}
	return table[len(prices)-1][0]
}
