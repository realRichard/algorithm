package list

// 力扣第 21 题
// 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/description/

// https://labuladong.github.io/algo/1/4/#合并两个有序链表

// 1 使用双指针，一个指向新链表的头节点，一个指向新链表的尾节点
// 2 注意需要特殊处理一开始头节点为空的情况
func MergeTwoListV1(list1 *ListNode, list2 *ListNode) *ListNode {
	// 新链表头节点
	header := (*ListNode)(nil)
	// 新链表尾节点
	tailer := (*ListNode)(nil)
	for list1 != nil || list2 != nil {
		// list1 为空，list2 不为空
		if list1 == nil {
			if header == nil {
				header = list2
				break
			} else {
				tailer.Next = list2
				break
			}
		}
		// list2 为空，list1 不为空
		if list2 == nil {
			if header == nil {
				header = list1
				break
			} else {
				tailer.Next = list1
				break
			}
		}
		// // 有点像归并算法后面的合并部分
		if list1.Val <= list2.Val {
			if header == nil {
				header = list1
				tailer = header
				list1 = list1.Next
			} else {
				tailer.Next = list1
				tailer = list1
				list1 = list1.Next
			}
		} else {
			if header == nil {
				header = list2
				tailer = header
				list2 = list2.Next
			} else {
				tailer.Next = list2
				tailer = list2
				list2 = list2.Next
			}
		}
	}
	return header
}

// 采用虚拟头节点，避免处理边界（一开始头节点为空的情况），降低代码复杂度
func MergeTwoListV2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 虚拟头节点
	dummy := new(ListNode)
	// 新链表尾节点
	p := dummy
	p1, p2 := list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
		// 不断向前进
		p = p.Next
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}
	return dummy.Next
}
