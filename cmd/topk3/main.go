package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(nums []int, l, r int) int {
	pivot := nums[r]
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

func randomizedPartition(nums []int, l, r int) int {
	i := l + rand.Intn(r-l+1)
	nums[r], nums[i] = nums[i], nums[r]
	return partition(nums, l, r)
}

func randomizedSelected(arr []int, l, r, k int) {
	pos := randomizedPartition(arr, l, r)
	num := pos - l + 1
	if k < num {
		randomizedSelected(arr, l, pos-1, k)
	} else if k > num {
		randomizedSelected(arr, pos+1, r, k-num)
	}
}

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	randomizedSelected(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func main() {
	arr := []int{
		1, 2, 3, 4, 5, 6, 7, 87, 8, 92, 3, 4, 23, 4, 56, 6, 568, 8, 67859,
	}
	k := 3
	fmt.Println(smallestK(arr, k))
}
