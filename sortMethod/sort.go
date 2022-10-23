package sortMethod

// 选择排序
// 1 循环每一个元素
// 2 与后面未排序的所有元素比较，找到未排序中最小的那个元素，交换位置
// 3 直到所有元素均排序完毕
// O(n²)
func SelectionSortV1(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			// 一次循环中可能要做多次交换
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func SelectionSortV2(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		// 先找到本轮循环中最小的元素
		min := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// 在做交换
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
}

// 插入排序，打扑克
// 将第一个元素看成已排序序列，循环后面未排序的所有元素
// 从后往前，在已排序序列中找到合适的位置
// 从后往前，如果小于已排序元素，则已排序元素往后移动一位
// 替换
// O(n²)
func InsertionSort(arr []int) {
	length := len(arr)
	for i := 1; i < length; i++ {
		tmp := arr[i]
		j := i
		for ; j > 0 && arr[j-1] > tmp; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = tmp
	}
}

// 冒泡排序
// O(n²)
func BubbleSort(arr []int) {
	length := len(arr)
	for i := 0; i < length-1; i++ {
		flag := false
		for j := 0; j < length-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			return
		}
	}
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}

// 归并排序
// O(nlogn)
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	return merge(left, right)
}

func partition(arr []int, lower, higher int) int {
	pivotIndex := higher
	pivot := arr[pivotIndex]
	higher--
	for lower < higher {
		for lower < higher && arr[lower] < pivot {
			lower++
		}
		for lower < higher && arr[higher] > pivot {
			higher--
		}
		if lower < higher && arr[lower] > arr[higher] {
			arr[lower], arr[higher] = arr[higher], arr[lower]
		}
	}
	if arr[higher] > arr[pivotIndex] {
		arr[higher], arr[pivotIndex] = arr[pivotIndex], arr[higher]
	}
	// arr[higher], arr[pivotIndex] = arr[pivotIndex], arr[higher]
	return higher
}

// 快速排序
// 最好 O(nlogn)
// 最坏 O(n²)
// https://mojotv.cn/algorithm/golang-quick-sort
func QuickSort(arr []int, lower, higher int) {
	if lower < higher {
		pivotIndex := partition(arr, lower, higher)
		QuickSort(arr, lower, pivotIndex-1)
		QuickSort(arr, pivotIndex+1, higher)
	}
}
