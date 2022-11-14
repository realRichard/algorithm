package list

// 力扣第 141 题
// 环形链表

// https://leetcode.cn/problems/linked-list-cycle/description/

// 判断链表是否包含环
// https://labuladong.github.io/algo/1/4/#判断链表是否包含环

// 1 直观的解法可以是利用 hashtable 记录走过的节点，如果出现重复则说明有环，不过有空间复杂度

// 2 利用快慢指针，如果快指针在走到结尾前，两个指针能相遇，说明快指针超过慢指针一圈了，说明有环
func HasCycle(header *ListNode) bool {
	fast, slow := header, header
	// 当 fast 走到末尾时
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// 力扣第 142 题
// 环形链表 II

// https://leetcode.cn/problems/linked-list-cycle-ii/description/

// 如果有环，如何计算找到这个环的起点
// 如果有环，那么快慢指针能够相遇，在相遇的时候，假设慢指针走了 k 步，那么快指针就走了 2k 步
// 这个时候假设相遇点距离环起点为 m
// 那么链表头节点距离环起点就是 k - m
// 随便把快慢指针中的一个指向链表头节点，然后两个指针按照相同的速度走 k - m 步就能在环起点相遇
func DetectCycle(header *ListNode) *ListNode {
	fast, slow := header, header
	// 判断是否有环
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	// 没有环直接退出
	if fast == nil || fast.Next == nil {
		return nil
	}
	// 重新指向头节点
	fast = header
	// 快慢指向同步前进，相遇点就是环起点
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
