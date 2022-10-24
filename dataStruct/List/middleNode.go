package list

// 力扣第 876 题
// 单链表的中点
// https://leetcode.cn/problems/middle-of-the-linked-list/description/

// https://labuladong.github.io/algo/1/4/#单链表的中点

// 1 直观的方法，先遍历链表算出节点数 n，在一次遍历找到中间节点

// 2 使用快慢指针，只用遍历一次

// 如果有两个中间节点返回第二个
func MiddleNode(header *ListNode) *ListNode {
	// 快慢指针初始化指向 header
	fast, slow := header, header
	// 快指针指向末尾时停止
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 慢指针指向中点
	return slow
}
