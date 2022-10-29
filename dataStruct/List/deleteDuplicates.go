package list

// 力扣第 83 题
// 删除排序链表中的重复元素
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for ; fast != nil; fast = fast.Next {
		if fast.Val != slow.Val {
			slow.Next = fast
			// 慢指针往前移
			slow = fast
		}
	}
	// 断开与后面重复元素的连接
	slow.Next = nil
	return head
}
