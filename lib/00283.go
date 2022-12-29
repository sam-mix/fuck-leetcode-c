package lib

func moveZeroes(nums []int) {
	// h: head 前面0的位置， t: 后面非零的位置，n: 需要操作的长度
	h, t, n := 0, 0, len(nums)-1
	for t <= n {
		if nums[n] == 0 {
			n--
			continue
		}
		if nums[t] != 0 {
			if t != h {
				nums[h], nums[t] = nums[t], nums[h]
			}
			h++
		}
		t++
	}
}
