package lib

func minOperations(nums []int) (sum int) {
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > cur {
			cur = nums[i]
			continue
		}
		cur = cur + 1
		sum += cur - nums[i]
	}
	return sum
}
