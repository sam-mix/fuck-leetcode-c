package lib

func longestPalindrome(s string) string {
	l := len(s)
	n := l/2 + 2
	mStr, mLen := "", 0
	for i := range s {
		lfrt := 0 // 0 左右都不靠；靠左 为1； 靠右 为2；左右都靠为3
		if i-1 >= 0 && s[i] == s[i-1] {
			lfrt = 1
		}
		if i+1 < l && s[i] == s[i+1] {
			if lfrt == 0 {
				lfrt = 2
			} else {
				lfrt = 3
			}
		}
		for j := 0; j < n; j++ {
			iMin, iMax := i-j, i+j
			if lfrt == 1 {
				iMax--
			}
			if lfrt == 2 {
				iMin++
			}
			// 超过上下边界直接结束
			if iMin < 0 || iMax >= l || s[iMin] != s[iMax] {
				if lfrt != 3 {
					break
				} else {
					iMax--
					if iMin < 0 || iMax >= l || s[iMin] != s[iMax] {
						break
					}
				}
			}
			curLen := iMax - iMin + 1
			if curLen > mLen {
				mLen = curLen
				mStr = s[iMin : iMax+1]
			}
		}
	}

	return mStr
}
