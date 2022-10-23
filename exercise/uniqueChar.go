package exercise

import (
	"strings"
)

// https://github.com/lifei6671/interview-go/blob/master/question/q002.md

// 问题描述

// 请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。 给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。 保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

// 解题思路

// 这里有几个重点，第一个是ASCII字符，ASCII字符字符一共有256个，其中128个是常用字符，可以在键盘上输入。128之后的是键盘上无法找到的。

// 然后是全部不同，也就是字符串中的字符没有重复的，再次，不准使用额外的储存结构，且字符串小于等于3000。

// 如果允许其他额外储存结构，这个题目很好做。如果不允许的话，可以使用golang内置的方式实现。

func IsUniqueChar1(s string) bool {
	for _, i := range s {
		if strings.Count(s, string(i)) > 1 {
			return false
		}
	}
	return true
}

func IsUniqueChar2(s string) bool {
	for k, i := range s {
		if strings.LastIndex(s, string(i)) != k {
			return false
		}
	}
	return true
}

// 思考
// 1 一个 uint64 类型可以保存 64 个不重复的 2 的次方数字
// 2 & 操作判断是否重复
// 3 | 操作用相应的 2 的次方数字把 uint64 里面相应的位给占了
func IsUniqString3(s string) bool {
	if len(s) == 0 || len(s) > 3000 {
		return false
	}
	// 256 个字符 256 = 64 + 64 + 64 + 64
	var mark1, mark2, mark3, mark4 uint64
	var mark *uint64
	for _, r := range s {
		n := uint64(r)
		if n < 64 {
			mark = &mark1
		} else if n < 128 {
			mark = &mark2
			n -= 64
		} else if n < 192 {
			mark = &mark3
			n -= 128
		} else {
			mark = &mark4
			n -= 192
		}
		if (*mark)&(1<<n) != 0 {
			return false
		}
		*mark = (*mark) | uint64(1<<n)
	}
	return true
}

// 源码解析

// 以上三种方法都可以实现这个算法。

// 第一个方法使用的是golang内置方法strings.Count,可以用来判断在一个字符串中包含的另外一个字符串的数量。

// 第二个方法使用的是golang内置方法strings.Index和strings.LastIndex，用来判断指定字符串在另外一个字符串的索引未知，分别是第一次发现位置和最后发现位置。

// 第三个方法使用的是位运算来判断是否重复，时间复杂度为o(n)，相比前两个方法时间复杂度低
