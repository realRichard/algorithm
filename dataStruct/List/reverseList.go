package list

// 力扣第 206 题
// 反转链表

// https://leetcode.cn/problems/reverse-linked-list/description/

func ReverseListV1(head *ListNode) *ListNode {
	var newHead *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = newHead
		newHead = tmp
	}
	return newHead
}

// https://labuladong.github.io/algo/2/19/19/#一递归反转整个链表

func ReverseListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	last := ReverseListV2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
