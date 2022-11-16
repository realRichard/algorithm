package list

// 力扣第 25 题
// K 个一组翻转链表

// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/

// https://labuladong.github.io/algo/2/19/20/#二代码实现

func ReverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 区间 [a, b) 包含 k 个待反转元素
	a, b := head, head
	i := k
	for ; i > 0 && b != nil; i-- {
		b = b.Next
	}
	// 不足 k 个，不需要反转，base case
	if i > 0 {
		return a
	}
	newHead := Reverse(a, b)
	foo := ReverseKGroup(b, k)
	a.Next = foo
	return newHead
}

// 反转区间 [a, b) 的元素，注意是左闭右开
func Reverse(a *ListNode, b *ListNode) *ListNode {
	pre, cur, nex := (*ListNode)(nil), a, a
	for cur != b {
		nex = nex.Next
		cur.Next = pre
		pre = cur
		cur = nex
	}
	return pre
}
