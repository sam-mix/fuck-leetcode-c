package lib

func beautySumMy(s string) (sum int) {
	if len(s) < 3 {
		return 0
	}
	for i := 3; i <= len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			m := make(map[byte]int)
			for k := j; k < j+i; k++ {
				sk := s[k]
				if n, ok := m[sk]; ok {
					m[sk] = n + 1
				} else {
					m[sk] = 1
				}
			}
			if len(m) < 2 {
				continue
			}
			max, min := -1, -1
			for _, v := range m {
				if max < 0 {
					max = v
				}
				if min < 0 {
					min = v
				}
				if v > max {
					max = v
				}
				if v < min {
					min = v
				}
			}
			sum += max - min
		}
	}
	return sum
}

func beautySum(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		maxFreq := 0
		for _, c := range s[i:] {
			cnt[c-'a']++
			maxFreq = max(maxFreq, cnt[c-'a'])
			minFreq := len(s)
			for _, c := range cnt {
				if c > 0 {
					minFreq = min(minFreq, c)
				}
			}
			ans += maxFreq - minFreq
		}
	}
	return
}
