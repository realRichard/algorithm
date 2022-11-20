package list

// 力扣第 234 题
// 回文链表

// https://leetcode.cn/problems/palindrome-linked-list/description/

// https://labuladong.github.io/algo/2/19/21/#一判断回文单链表

// 1 回文链表，结构对称，可以反转生成新链表进行判断

// 2 可以把节点值收集到数组里面，然后利用数组双指针进行判断

// 左指针
var left *ListNode

// 保存结果
var res bool

func IsPalindromeV1(head *ListNode) bool {
	res = false
	left = head
	return isPalindrome(head)
}

// 3 利用递归后序遍历，模仿双指针进行判断
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	result := isPalindrome(head.Next)
	res = result && (head.Val == left.Val)
	left = left.Next
	return res
}

// 4 利用快慢指针找到链表中间节点，反转链表后半段，然后进行判断
func IsPalindromeV2(head *ListNode) bool {
	slow, fast := head, head
	if fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 注意细节处理，fast != nil，表示链表个数为奇数，此时 slow 指针还要在往前一步
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := ReverseListV1(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}
