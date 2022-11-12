package dp

// 力扣第 198 题
// 打家劫舍

// https://leetcode.cn/problems/house-robber/description/

// https://labuladong.github.io/algo/1/14/#打家劫舍-i
// https://mp.weixin.qq.com/s/z44hk0MW14_mAQd7988mfw

func RobV1(nums []int) int {
	return robV1(nums, 0)
}

func robV1(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	res := 0
	res1 := robV1(nums, start+1)
	res2 := nums[start] + robV1(nums, start+2)
	if res1 > res2 {
		res = res1
	} else {
		res = res2
	}
	return res
}

func RobV2(nums []int) int {
	table := make([]int, len(nums))
	for index := range table {
		table[index] = -1
	}
	return robV2(nums, 0, table)
}

func robV2(nums []int, start int, table []int) int {
	if start >= len(nums) {
		return 0
	}
	if table[start] != -1 {
		return table[start]
	}
	res := 0
	res1 := robV2(nums, start+1, table)
	res2 := nums[start] + robV2(nums, start+2, table)
	if res1 > res2 {
		res = res1
	} else {
		res = res2
	}
	table[start] = res
	return res
}

func RobV3(nums []int) int {
	// table[i] 表示从第 i 家开始偷，能偷到的最大金额
	table := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		if table[i+1] > table[i+2]+nums[i] {
			table[i] = table[i+1]
		} else {
			table[i] = table[i+2] + nums[i]
		}
	}
	return table[0]
}

func RobV4(nums []int) int {
	dp_i_1, dp_i_2 := 0, 0
	for i := len(nums) - 1; i >= 0; i-- {
		tmp := dp_i_1
		if dp_i_1 > dp_i_2+nums[i] {
			dp_i_1 = dp_i_1
		} else {
			dp_i_1 = dp_i_2 + nums[i]
		}
		dp_i_2 = tmp
	}
	return dp_i_1
}
