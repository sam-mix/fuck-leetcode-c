package lib

func strStr(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return -1
	}
	i := 0
	for ; i <= n; i++ {
		if n-i < m {
			return -1
		}
		j := 0
		for ; j < m; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j >= m {
			break
		}
	}
	return i
}
