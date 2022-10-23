package find

func BinarySearch(arr []int, target int) int {
	lower := 0
	higher := len(arr) - 1
	for lower <= higher {
		middle := (lower + higher) / 2
		if arr[middle] == target {
			return middle
		} else if arr[middle] < target {
			lower++
		} else {
			higher--
		}
	}
	return -1
}
