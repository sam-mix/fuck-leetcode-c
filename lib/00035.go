package lib

func searchInsert(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	for left <= right {
		if nums[right] < target {
			return right + 1
		}
		if nums[left] > target {
			return left
		}
		mid = (right + left) / 2
		if target == nums[mid] {
			return mid
		}
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return mid
}
