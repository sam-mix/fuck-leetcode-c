package lib

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l <= 1 {
		return strs[0]
	}
	n := len(strs[0])
	if n <= 0 {
		return ""
	}
	first := strs[0]
	common := make([]byte, 0)
	for i := 0; i < n; i++ {
		b := first[i]
		for j := 1; j < l; j++ {
			str := strs[j]
			if len(str) <= i || str[i] != b {
				return string(common)
			}
		}
		common = append(common, b)
	}

	return string(common)
}
