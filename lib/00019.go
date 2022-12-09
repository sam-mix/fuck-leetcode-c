package lib

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := 0
	p := head
	for p != nil {
		l++
		p = p.Next
	}
	if l == 1 && n == 1 {
		return nil
	}
	if l == n {
		head = head.Next
		return head
	}

	m := l - n - 1

	i := 0
	p = head
	for p != nil {
		if i == m && p.Next != nil {
			p.Next = p.Next.Next
		}
		i++
		p = p.Next
	}
	return head
}
