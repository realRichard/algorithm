package dp

// 力扣第 322 题
// 零钱兑换

// https://leetcode.cn/problems/coin-change/description/

// https://labuladong.github.io/algo/1/7/#二凑零钱问题

// 假设你有面值为 1, 2, 5 的硬币，你想求 amount = 11 时的最少硬币数（原问题），如果你知道凑出 amount = 10, 9, 6 的最少硬币数（子问题），你只需要把子问题的答案加一（再选一枚面值为 1, 2, 5 的硬币），求个最小值，就是原问题的答案

// 定义：要凑出金额 n，至少要 dp(coins, n) 个硬币
// 递归算法的时间复杂度分析：子问题总数 x 解决每个子问题所需的时间
// 子问题总数为递归树的节点个数，但算法会进行剪枝，剪枝的时机和题目给定的具体硬币面额有关，所以可以想象，这棵树生长的并不规则，确切算出树上有多少节点是比较困难的。对于这种情况，我们一般的做法是按照最坏的情况估算一个时间复杂度的上界。
// 假设目标金额为 n，给定的硬币个数为 k，那么递归树最坏情况下高度为 n（全用面额为 1 的硬币），然后再假设这是一棵满 k 叉树，则节点的总数在 k^n 这个数量级。
// 接下来看每个子问题的复杂度，由于每次递归包含一个 for 循环，复杂度为 O(k)，相乘得到总时间复杂度为 O(k^n)，指数级别。
func CoinChangeV1(coins []int, amount int) int {
	// base case
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	n := 2 << 32
	// 不同的选择，导致状态发生转换
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := CoinChangeV1(coins, amount-coin)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		if subProblem+1 < n {
			n = subProblem + 1
		}
	}
	if n == 2<<32 {
		return -1
	}
	return n
}

func dp(coins []int, amount int, table []int) int {
	// base case
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	if table[amount] != -2 {
		return table[amount]
	}
	n := 2 << 32
	for _, coin := range coins {
		// 计算子问题的结果
		subProblem := dp(coins, amount-coin, table)
		// 子问题无解则跳过
		if subProblem == -1 {
			continue
		}
		// 在子问题中选择最优解，然后加一
		if subProblem+1 < n {
			n = subProblem + 1
		}
	}
	if n == 2<<32 {
		table[amount] = -1
		return -1
	}
	table[amount] = n
	return n
}

// 很显然「备忘录」大大减小了子问题数目，完全消除了子问题的冗余，所以子问题总数不会超过金额数 n，即子问题数目为 O(n)。处理一个子问题的时间不变，仍是 O(k)，所以总的时间复杂度是 O(kn)。
func CoinChangeV2(coins []int, amount int) int {
	table := make([]int, amount+1)
	for index := range table {
		table[index] = -2
	}
	return dp(coins, amount, table)
}

func CoinChangeV3(coins []int, amount int) int {
	// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。
	table := make([]int, amount+1)
	// 为啥 dp 数组中的值都初始化为 amount + 1 呢，因为凑成 amount 金额的硬币数最多只可能等于 amount（全用 1 元面值的硬币），所以初始化为 amount + 1 就相当于初始化为正无穷，便于后续取最小值
	// 为啥不直接初始化为 int 型的最大值 Integer.MAX_VALUE 呢？因为后面有 dp[i - coin] + 1，这就会导致整型溢出
	for index := range table {
		table[index] = amount + 1
	}
	table[0] = 0
	// 外层 for 循环在遍历所有状态的所有取值
	for i := 0; i <= amount; i++ {
		// 内层 for 循环在求所有选择的最小值
		for _, coin := range coins {
			// 子问题无解，跳过
			if i-coin < 0 {
				continue
			}
			if table[i] > table[i-coin]+1 {
				table[i] = table[i-coin] + 1
			}
		}
	}
	if table[amount] == amount+1 {
		return -1
	} else {
		return table[amount]
	}
}
