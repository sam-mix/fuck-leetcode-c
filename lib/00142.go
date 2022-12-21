package lib

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func detectCycle(head *ListNode) *ListNode {
// 	m := make(map[*ListNode]struct{})
// 	for head != nil && head.Next != nil {
// 		m[head] = struct{}{}
// 		if _, ok := m[head.Next]; ok {
// 			return head.Next
// 		}
// 		head = head.Next
// 	}
// 	return nil
// }

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for fast != slow {
		fast, slow = fast.Next, slow.Next
	}
	return fast
}
