package array

import "sort"

// 力扣第 18 题
// 四数之和

// https://leetcode.cn/problems/4sum/description/

func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		foo := target - nums[i]
		threeSumResult := ThreeSumTarget(nums[i+1:], foo)
		for index, item := range threeSumResult {
			if len(item) != 0 {
				threeSumResult[index] = append(threeSumResult[index], nums[i])
				res = append(res, threeSumResult[index])
			}
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
