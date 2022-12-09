package main

import (
	"container/heap"
	"fmt"
)

type heapTest struct {
	name string
	age  int
}

type TestHeap []heapTest // 定义一个类型

func (h TestHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h TestHeap) Less(i, j int) bool { // 绑定less方法
	return h[i].age < h[j].age // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h TestHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *TestHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *TestHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(heapTest))
}

func main() {
	h := &TestHeap{
		{"a", 5},
		{"b", 4},
		{"c", 3},
		{"d", 2},
		{"e", 1},
		{"f", 0},
	} // 创建slice
	heap.Init(h) // 初始化heap
	// fmt.Println(heap.Pop(h)) // 调用pop
	fmt.Println(*h)
	heap.Push(h, heapTest{"z", -1})
	for len(*h) > 0 {
		ht := heap.Pop(h).(heapTest)
		fmt.Printf("%s => %d\t", ht.name, ht.age)
	}
	fmt.Println()

}
