package array

// 力扣第 344 题
// 反转字符串

// https://leetcode.cn/problems/reverse-string/description/

func ReverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
