package exercise

import "sort"

// method 1 穷举法，从数组中任意找两个数，看其和是否=sum。时间复杂度O(N^2)

// method 2 先排序，前后指针.  时间复杂度：快排O(NlogN)，查找O(N)；所以总的时间复杂度为：O(NlogN)
func TwoSum(arr []int, sum int) bool {
	length := len(arr)
	if length < 2 {
		return false
	}
	sort.Ints(arr)
	lower := 0
	higher := length - 1
	for lower < higher {
		result := arr[lower] + arr[higher]
		if result == sum {
			return true
		} else if result < sum {
			lower++
		} else if result > sum {
			higher--
		}
	}
	return false
}

// method 3 hashtable  但需要空间换时间，空间复杂度为O(n)
