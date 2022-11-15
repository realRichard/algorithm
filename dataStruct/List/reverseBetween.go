package list

// 力扣第 92 题
// 反转链表 II

// https://leetcode.cn/problems/reverse-linked-list-ii/description/

// https://labuladong.github.io/algo/2/19/19/#三反转链表的一部分

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return ReverseN(head, right)
	}
	// 前进到反转的起点触发 base case
	head.Next = ReverseBetween(head.Next, left-1, right-1)
	return head
}

// 后驱节点
var successor *ListNode

// 反转以 head 为起点的 n 个节点，返回新的头结点
func ReverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		//  记录第 n + 1 个节点
		successor = head.Next
		return head
	}
	// // 以 head.next 为起点，需要反转前 n - 1 个节点
	last := ReverseN(head.Next, n-1)
	head.Next.Next = head
	// 让反转之后的 head 节点和后面的节点连起来
	head.Next = successor
	return last
}
