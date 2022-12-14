package lib

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	cur, pre := head, head
	for cur != nil {
		if cur != pre && cur.Val == pre.Val {
			pre.Next = cur.Next
			cur = pre.Next
		} else {
			pre = cur
			cur = cur.Next
		}
	}
	return head
}
