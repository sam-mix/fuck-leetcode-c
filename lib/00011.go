package lib

// 双层循环 （无脑写）
// 超时
func maxAreaTimeOut(height []int) int {
	hs := height
	max := 0
	n := len(hs) - 1
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			mh := hs[i]
			if hs[j] < hs[i] {
				mh = hs[j]
			}
			if mh <= 0 {
				continue
			}
			v := (j - i) * mh
			if v > max {
				max = v
			}
		}
	}

	return max
}

// 左右坐标向中间移动， 矮的一方移动， 如果同样高就同时移动， 因为只移动一方体积必然都会变少：最小高度不变或变小，宽度一定变小。
func maxArea(hights []int) int {
	hs := hights
	i, j := 0, len(hs)-1
	max, left, right, mh, cur := 0, 0, 0, 0, 0
	for i < j {
		left = hs[i]
		right = hs[j]
		mh = left
		if mh > right {
			mh = right
		}
		cur = mh * (j - i)
		if cur > max {
			max = cur
		}
		if right < left {
			j--
		} else if left < right {
			i++
		} else {
			i++
			j--
		}

	}
	return max
}
