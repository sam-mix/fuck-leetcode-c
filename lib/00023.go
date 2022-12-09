package lib

import (
	"container/heap"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil {
		return nil
	}
	var head *ListNode = nil
	for _, v := range lists {
		head = mergeTwoLists(head, v)
	}
	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Time: O(n*log(k)), Space: O(k) n是总节点数量，k是链表个数
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil // 链表为空或长度为0，直接返回空指针
	}
	var h IntHeap // 最小堆用于维护当前k个节点
	heap.Init(&h) // 用于节点间的比较

	for _, list := range lists {
		// 数组中非空的链表加入到最小堆
		if list != nil {
			heap.Push(&h, list)
		}
	}
	// 定义dummy节点用于统一处理
	dummy := &ListNode{}
	p := dummy // p初始指向dummy节点

	// 当最小堆不为空时，不断执行以下操作
	for h.Len() > 0 { // 取出堆顶元素，即取出最小值节点
		min := heap.Pop(&h).(*ListNode)
		fmt.Println(min.Val)
		p.Next = min // 游标p指向最小值节点
		p = p.Next   // p向后移动一个位置
		// 这样就确定一个节点在合并链表中的位置
		if min.Next != nil { // 如果最小值节点后面的节点非空
			heap.Push(&h, min.Next) // 则把最小值节点后面的节点加入到最小堆中
		}
	}
	return dummy.Next // 最后只要返回dummy.Next即可
}

type IntHeap []*ListNode

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
