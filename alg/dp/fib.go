package dp

// 力扣第 509 题
// 斐波那契数

// https://leetcode.cn/problems/fibonacci-number/description/

// https://labuladong.github.io/algo/1/7/#一斐波那契数列

// 递归算法的时间复杂度怎么计算？就是用子问题个数乘以解决一个子问题需要的时间

// 暴力递归
// 首先计算子问题个数，即递归树中节点的总数。显然二叉树节点总数为指数级别，所以子问题个数为 O(2^n)。
// 然后计算解决一个子问题的时间，在本算法中，没有循环，只有 f(n - 1) + f(n - 2) 一个加法操作，时间为 O(1)。
// 所以，这个算法的时间复杂度为二者相乘，即 O(2^n)，指数级别，爆炸。
func FibV1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return FibV1(n-1) + FibV1(n-2)
}

func helper(table []int, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if table[n] != 0 {
		return table[n]
	}
	table[n] = helper(table, n-1) + helper(table, n-2)
	return table[n]
}

// 带备忘录的递归解法
// 子问题个数，即图中节点的总数，由于本算法不存在冗余计算，子问题就是 f(1), f(2), f(3) … f(20)，数量和输入规模 n = 20 成正比，所以子问题个数为 O(n)
// 解决一个子问题的时间，同上，没有什么循环，时间为 O(1)
// 所以，本算法的时间复杂度是 O(n)，比起暴力算法，是降维打击
func FibV2(n int) int {
	table := make([]int, n+1)
	return helper(table, n)
}

func FibV3(n int) int {
	table := make([]int, 2, n+1)
	table[0] = 0
	table[1] = 1
	for i := 2; i <= n; i++ {
		table = append(table, table[i-1]+table[i-2])
	}

	return table[n]
}

// 根据斐波那契数列的状态转移方程，当前状态只和之前的两个状态有关，其实并不需要那么长的一个 DP table 来存储所有的状态，只要想办法存储之前的两个状态就行了
// 空间复杂度降为 O(1)
// 这一般是动态规划问题的最后一步优化，如果我们发现每次状态转移只需要 DP table 中的一部分，那么可以尝试缩小 DP table 的大小，只记录必要的数据，从而降低空间复杂度
func FibV4(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	dpI1 := 0
	dpI2 := 1
	dpI := 0
	for i := 2; i <= n; i++ {
		dpI = dpI1 + dpI2
		dpI1 = dpI2
		dpI2 = dpI
	}
	return dpI
}
