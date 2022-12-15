package lib

import (
	"sort"
)

func merge4(nums1 []int, m int, nums2 []int, n int) {
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func merge3(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

func merge(a []int, m int, b []int, n int) {
	sorted, p1, p2 := make([]int, 0, m+n), 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, b[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, a[p1:]...)
			break
		}
		if a[p1] < b[p2] {
			sorted = append(sorted, a[p1])
			p1++
		} else {
			sorted = append(sorted, b[p2])
			p2++
		}
	}

}
