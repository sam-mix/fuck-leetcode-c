package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	n := len(arr)
	if k > n {
		sort.Ints(arr)
		return arr
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 87, 8, 92, 3, 4, 23, 4, 56, 6, 568, 8, 67859,
	}
	k := 100
	fmt.Println(smallestK(arr, k))
}
