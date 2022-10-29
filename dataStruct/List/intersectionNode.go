package list

// 力扣第 160 题
// 两个链表是否相交
// https://leetcode.cn/problems/intersection-of-two-linked-lists/

// https://labuladong.github.io/algo/1/4/#两个链表是否相交

// 1 直接的解法可能是先遍历一条链表，用 hashtable 将节点都存起来
// 然后在遍历另外一条链表，判断是否有相同的节点，不过需要额外的空间

// 2 难点在于，由于两条链表的长度可能不同，两条链表之间的节点无法对应
// 如果用两个指针 p1 和 p2 分别在两条链表上前进，并不能同时走到公共节点，也就无法得到相交节点 c1
// 解决这个问题的关键是，通过某些方式，让 p1 和 p2 能够同时到达相交节点 c1
// 所以，我们可以让 p1 遍历完链表 A 之后开始遍历链表 B，让 p2 遍历完链表 B 之后开始遍历链表 A，这样相当于「逻辑上」两条链表接在了一起
// 如果这样进行拼接，就可以让 p1 和 p2 同时进入公共部分，也就是同时到达相交节点 c1
// 空间 O(1)，时间 O(n)
func GetIntersectionNodeV1(head1, head2 *ListNode) *ListNode {
	p1, p2 := head1, head2
	for p1 != p2 {
		if p1 == nil {
			p1 = head2
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = head1
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

// 3 如果把两条链表首尾相连，那么寻找两条链表的交点，就转化为寻找环起点的问题了
// 改变了原链表的结构，解题完后，需要把链表结构改回来，不然可能无法通过

// 4 寻找两条链表的核心在于能够让 p1 p2 指针同时到达相交节点，那么可以通过预先遍历计算出两条链表的长度
func GetIntersectionNodeV2(head1, head2 *ListNode) *ListNode {
	len1 := 0
	for head1 := head1; head1 != nil; head1 = head1.Next {
		len1++
	}
	len2 := 0
	for head2 := head2; head2 != nil; head2 = head2.Next {
		len2++
	}
	p1 := head1
	p2 := head2
	if len1 < len2 {
		diff := len2 - len1
		for i := diff; i > 0; i-- {
			p2 = p2.Next
		}
	} else if len1 > len2 {
		diff := len1 - len2
		for i := diff; i > 0; i-- {
			p1 = p1.Next
		}
	}
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}
