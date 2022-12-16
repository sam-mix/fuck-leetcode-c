package lib

func minElements(nums []int, limit int, goal int) int {
	if limit <= 0 {
		return 0
	}
	sum := -goal
	for _, v := range nums {
		if (v < 0 && -v > limit) || v > limit {
			return 0
		}
		sum += v
	}
	if sum < 0 {
		sum = -sum
	}
	ans := sum / limit
	if sum%limit > 0 {
		ans++
	}
	return ans
}
