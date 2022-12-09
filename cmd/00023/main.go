package main

import "sam-mix.ltd/fuck-leetcode-c/lib"

func main() {
	list := []*lib.ListNode{
		{Val: 6, Next: &lib.ListNode{Val: 12, Next: nil}},
		{Val: 5, Next: &lib.ListNode{Val: 11, Next: nil}},
		{Val: 4, Next: &lib.ListNode{Val: 10, Next: nil}},
		{Val: 3, Next: &lib.ListNode{Val: 9, Next: nil}},
		{Val: 2, Next: &lib.ListNode{Val: 8, Next: nil}},
		{Val: 1, Next: &lib.ListNode{Val: 7, Next: nil}},
	}

	lib.MergeKLists2(list)
}
