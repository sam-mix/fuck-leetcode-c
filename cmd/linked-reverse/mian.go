package main

import "fmt"

type listNode struct {
	val  interface{}
	next *listNode
}

func printList(head *listNode) {
	for head != nil {
		fmt.Println(head.val)
		head = head.next
	}
}

func main() {
	head := &listNode{
		1,
		&listNode{
			2,
			&listNode{
				3,
				&listNode{
					4,
					nil,
				},
			},
		},
	}
	fmt.Println("before: ")
	printList(head)
	after := reverse(head)
	fmt.Println("after: ")
	printList(after)
}

func reverse(head *listNode) (newHead *listNode) {
	if head == nil || head.next == nil {
		return head
	}
	next := head.next
	head.next = nil
	newHead = reverse(next)
	next.next = head
	return
}

// func reverse(head *listNode) (newHead *listNode) {
// 	for head != nil {
// 		next := head.next
// 		head.next = newHead
// 		newHead = head
// 		head = next
// 	}
// 	return
// }
