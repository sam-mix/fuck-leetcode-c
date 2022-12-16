package lib

func generate(n int) (ans [][]int) {
	for i := 0; i < n; n++ {
		ansi := []int{1}
		for j := 1; j < i; j++ {
			ansi[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		if i > 0 {
			ansi[i] = 1
		}
		ans[i] = ansi
	}
	return
}
