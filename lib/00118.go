package lib

func generate(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ansi := make([]int, i+1)
		ansi[0] = 1
		for j := 1; j < i; j++ {
			ansi[j] = ans[i-1][j-1] + ans[i-1][j]
		}
		if i > 0 {
			ansi[i] = 1
		}
		ans[i] = ansi
	}
	return ans
}
