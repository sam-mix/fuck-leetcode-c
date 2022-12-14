package lib

func climbStairs(n int) int {
	step1, step2 := 1, 2
	for i := 1; i < n; i++ {
		step2, step1 = step1+step2, step2
	}
	return step1
}
