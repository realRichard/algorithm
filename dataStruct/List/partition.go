package list

// 力扣第 86 题
// 单链表的分解
// https://leetcode.cn/problems/partition-list/description/

// https://labuladong.github.io/algo/1/4/#单链表的分解

func Partition(header *ListNode, x int) *ListNode {
	// 比 x 小的新链表的虚拟头节点
	dummy1 := new(ListNode)
	// 比 x 小的新链表的尾节点
	p1 := dummy1
	// 保存比 x 大的链表的虚拟头节点
	dummy2 := new(ListNode)
	// 比 x 大的链表的尾节点
	p2 := dummy2
	p := header
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		// 断开原链表中每个节点的 next 指针
		tmp := p.Next
		p.Next = nil
		p = tmp
	}
	// // 一定要把较大的链表的最后一个元素的 next 指向 nil，不然可能成环
	// p2.Next = nil
	p1.Next = dummy2.Next
	return dummy1.Next
}
