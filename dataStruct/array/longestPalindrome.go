package array

// 回文串判断
// 首先明确一下，回文串就是正着读和反着读都一样的字符串。
// 比如说字符串 aba 和 abba 都是回文串，因为它们对称，反过来还是和本身一样；反之，字符串 abac 就不是回文串

func IsPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 力扣第 5 题
// 最长回文子串

// https://leetcode.cn/problems/longest-palindromic-substring/

// 找回文串的难点在于，回文串的的长度可能是奇数也可能是偶数，解决该问题的核心是从中心向两端扩散的双指针技巧
// 如果回文串的长度为奇数，则它有一个中心字符；如果回文串的长度为偶数，则可以认为它有两个中心字符。所以我们可以先实现这样一个函数：
// 这样，如果输入相同的 l 和 r，就相当于寻找长度为奇数的回文串，如果输入相邻的 l 和 r，则相当于寻找长度为偶数的回文串
func palindrome(s string, left, right int) string {
	length := len(s)
	for left >= 0 && right < length && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}

func LongestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		// 以 s[i] 为中心的最长回文子串
		subStr := palindrome(s, i, i)
		if len(subStr) > len(result) {
			result = subStr
		}
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		subStr = palindrome(s, i, i+1)
		if len(subStr) > len(result) {
			result = subStr
		}
	}
	return result
}
