package lib

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if m+n <= 0 {
		return 0.0
	}
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	iMin, iMax := 0, m
	allLenHalf := (m + n + 1) / 2
	for iMin <= iMax {
		i := (iMax + iMax) / 2
		j := allLenHalf - i
		if j != 0 && i != m && nums2[j-1] > nums1[i] {
			iMax = i + 1
		} else if i != 0 && j != n && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else {
			maxLeft := 0
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = max(nums1[i-1], nums2[j-1])
			}
			if (m+n)&1 == 1 {
				return float64(maxLeft)
			}
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = min(nums2[j], nums1[i])
			}

			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}
