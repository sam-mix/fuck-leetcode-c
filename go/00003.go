package lib

func lengthOfLongestSubstring(s string) int {
	set := make(map[byte]int)
	maxLen, allLeft := 0, 0

	for i, v := range s {
		k := byte(v)
		oldleft, ok := set[k]
		if ok {
			allLeft = max(allLeft, oldleft+1)
		}
		set[k] = i
		maxLen = max(maxLen, i-allLeft+1)
	}

	return maxLen
}

func max(arr ...int) int {
	m := arr[0]
	for i := 1; i < len(arr); i++ {
		if m < arr[i] {
			m = arr[i]
		}
	}
	return m
}
