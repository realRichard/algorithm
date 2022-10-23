package list

func ReverseLinkedListV1(head *ListNode) *ListNode {

	var newHead *ListNode
	for head != nil {
		tmp := head
		head = head.Next
		tmp.Next = newHead
		newHead = tmp
	}
	return newHead
}
