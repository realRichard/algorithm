package list

// 单链表的倒数第 k 个节点

// https://labuladong.github.io/algo/1/4/#单链表的倒数第-k-个节点

// 1 通过计算得出倒数第 k 个节点就是 (n - k + 1), 但是需要先一趟循环算出 n, 然后在一趟循环找到倒数第 k 个节点

// 2 使用双指针，第一个指针比第二个指针先走 k 步
func KthNodeFromEnd(list *ListNode, k int) *ListNode {
	p1, p2 := list, list
	for p1 != nil && k > 0 {
		p1 = p1.Next
		k--
	}
	if k != 0 {
		return nil
	}
	for p1 != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p2
}

// 力扣第 19 题
// 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

func RemoveNthFromEnd(header *ListNode, n int) *ListNode {
	// 使用虚拟节点，防止 nil 空指针出现
	// 例如，链表总共有五个节点，要求删除倒数第五个节点，也就是第一个节点
	// 按照逻辑，我们应该先找到倒数第六个节点，但由于第一个节点前面已经没有节点了，就会出错
	// 但是使用虚拟节点，就避免了这个问题，并且可以正确的删除
	dummy := new(ListNode)
	dummy.Next = header
	// 要删除倒数第 n 个节点，先要找到倒数第 n + 1 的节点
	tmp := KthNodeFromEnd(dummy, n+1)
	if tmp == nil {
		return nil
	}
	tmp.Next = tmp.Next.Next
	return dummy.Next
}
