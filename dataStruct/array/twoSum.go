package array

// 力扣第 1 题
// 两数之和

// https://leetcode.cn/problems/two-sum/description/

func TwoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for index, item := range nums {
		foo := target - item
		if i, ok := table[foo]; ok {
			return []int{i, index}
		}
		table[item] = index
	}
	return []int{}
}
