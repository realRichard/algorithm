package array

import (
	"sort"
)

// nSum

// https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q

func NSum(nums []int, n int, target int) [][]int {
	// 要先排序， 不然在递归函数里面排序效率极低
	sort.Ints(nums)
	return nSum(nums, n, target, 0)
}

func nSum(nums []int, n int, target int, start int) [][]int {
	res := make([][]int, 0)
	// 至少是 2Sum，且数组大小不应该小于 n
	if n < 2 || n > len(nums) {
		return res
	}
	if n == 2 {
		// 双指针那一套操作
		left, right := start, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum > target {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else if sum == target {
				res = append(res, []int{nums[left], nums[right]})
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	} else {
		// n > 2 时，递归计算 (n-1)Sum 的结果
		for i := start; i < len(nums); i++ {
			foo := target - nums[i]
			result := nSum(nums, n-1, foo, i+1)
			for index, item := range result {
				if len(item) != 0 {
					result[index] = append(result[index], nums[i])
					res = append(res, result[index])
				}
			}
			for i < len(nums)-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return res
}
