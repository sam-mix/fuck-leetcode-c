package main

import (
	"container/heap"
	"errors"
	"fmt"
)

type HeapNode struct {
	value  int32 //值是什么
	arrNum int   //来自哪个数组
	index  int   //来自数组的哪个位置
}

type ArrayHeap []*HeapNode

func (arrayheap ArrayHeap) Len() int {
	return len(arrayheap)
}

func (arrayheap ArrayHeap) Less(i, j int) bool {
	return arrayheap[i].value > arrayheap[j].value //大根堆
}

func (arrayheap ArrayHeap) Swap(i, j int) {
	arrayheap[i], arrayheap[j] = arrayheap[j], arrayheap[i]
}

func (arrayheap *ArrayHeap) Push(node interface{}) {
	item := node.(*HeapNode)
	*arrayheap = append(*arrayheap, item)
}

func (arrayheap *ArrayHeap) Pop() interface{} {
	old := *arrayheap
	n := len(old)
	item := old[n-1]
	*arrayheap = old[0 : n-1]
	return item
}

func (arrayheap *ArrayHeap) IsEmpty() bool {
	return len(*arrayheap) == 0
}

func getMatriValue(matrix *[3][5]int32, i, j int) (value int32, err error) {
	if i < 0 || i > len(*matrix)-1 {
		return 0, errors.New("out of matrix index")
	}
	if j < 0 || j > len(matrix[0])-1 {
		return 0, errors.New("out of matrix index")
	}
	return matrix[i][j], nil
}

// func printTopK( int topK) {

// }

func main() {
	var matrix = [][]int32{
		{219, 405, 538, 845, 971},
		{148, 558, 0, 0, 0},
		{52, 99, 348, 691, 0},
	}

	myHeap := make(ArrayHeap, 0)
	heap.Init(&myHeap)

	for arri, arr := range matrix {
		for i, v := range arr {
			heap.Push(&myHeap, &HeapNode{
				value:  v,
				arrNum: arri,
				index:  i,
			})
		}
	}

	for i := 0; i < 5; i++ {
		item := heap.Pop(&myHeap).(*HeapNode)
		fmt.Printf("value = %d, arrNum = %d, index = %d\n", item.value, item.arrNum, item.index)

	}
}
