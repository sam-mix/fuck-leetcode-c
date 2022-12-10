package lib

import (
	"sort"
)

func maxHeight(cuboids [][]int) int {
	for _, v := range cuboids {
		sort.Ints(v)
	}
	sum := func(xs []int) int {
		s := 0
		for _, v := range xs {
			s += v
		}
		return s
	}
	sort.Slice(cuboids, func(i, j int) bool {
		return sum(cuboids[i]) <= sum(cuboids[j])
	})
	mf := func(xs ...int) int {
		m := xs[0]
		for i := 1; i < len(xs); i++ {
			if m < xs[i] {
				m = xs[i]
			}
		}
		return m
	}
	max := 0
	n := len(cuboids)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = cuboids[i][2]
		for j := 0; j < i; j++ {
			if cuboids[i][0] >= cuboids[j][0] && cuboids[i][1] >= cuboids[j][2] && cuboids[i][2] >= cuboids[j][2] {
				dp[i] = mf(dp[i], dp[i]+cuboids[j][2])
			}
		}

		max = mf(max, dp[i])
	}

	return max
}
