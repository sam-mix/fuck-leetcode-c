package lib

import (
	"sort"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	cases := []struct {
		in struct {
			a []int
			b []int
		}
	}{
		{
			in: struct {
				a []int
				b []int
			}{[]int{1, 2, 3}, []int{3, 3, 4, 5, 6, 7}},
		},
		{
			in: struct {
				a []int
				b []int
			}{[]int{1, 2, 3, 4, 5, 6, 7, 8}, []int{3, 3, 4, 5, 6, 7}},
		},
		{
			in: struct {
				a []int
				b []int
			}{[]int{1, 2, 3}, []int{3, 3, 3, 4, 4, 4, 4, 4, 4, 5, 6, 7}},
		},
		{
			in: struct {
				a []int
				b []int
			}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, []int{3, 3, 4, 5, 6, 7}},
		},
		{
			in: struct {
				a []int
				b []int
			}{[]int{1, 100}, []int{3, 3, 4, 5, 6, 7}},
		}, {
			in: struct {
				a []int
				b []int
			}{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1001}, []int{300, 400, 500, 600, 700, 800, 9000}},
		}, {
			in: struct {
				a []int
				b []int
			}{[]int{100}, []int{10, 100000, 999999}},
		}, {
			in: struct {
				a []int
				b []int
			}{[]int{1000, 10000, 100000}, []int{1, 3, 4, 5, 6, 7, 1001, 1100, 9999, 10002, 18888, 100001}},
		}, {
			in: struct {
				a []int
				b []int
			}{[]int{}, []int{}},
		}, {
			in: struct {
				a []int
				b []int
			}{[]int{1}, []int{}},
		},
		{
			in: struct {
				a []int
				b []int
			}{[]int{1}, []int{1}},
		},
	}
	for _, v := range cases {
		assertEqual(t, helpFindMedianSortedArrays(v.in), findMedianSortedArrays(v.in.a, v.in.b), v.in)
	}
}

func helpFindMedianSortedArrays(in struct {
	a []int
	b []int
}) float64 {
	a := append(in.a, in.b...)
	l := len(a)
	if l <= 0 {
		return 0.0
	}
	if l <= 1 {
		return float64(a[0])
	}
	sort.Ints(a)
	if l&1 == 1 {
		return float64(a[(l-1)/2])
	}
	m := l / 2
	return float64(a[m]+a[m-1]) / 2.0
}
