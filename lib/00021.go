package lib

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head, other := list1, list2
	if head.Val > list2.Val {
		head = list2
		other = list1
	}
	p := head
	for p != nil && other != nil {
		if p.Next == nil {
			p.Next = other
			break
		}
		if p.Next.Val <= other.Val {
			p = p.Next
			continue
		}
		tempNext := p.Next
		p.Next = other
		other = p.Next.Next
		p.Next.Next = tempNext
	}

	return head
}
