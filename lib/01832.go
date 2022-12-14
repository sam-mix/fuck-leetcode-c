package lib

func checkIfPangramMap(sentence string) bool {
	N := 25
	if len(sentence) <= N {
		return false
	}
	set := make(map[rune]struct{}, N)
	for _, v := range sentence {
		if _, ok := set[v]; !ok {
			if len(set) >= N {
				return true
			}
			set[v] = struct{}{}
		}
	}
	return false
}

func checkIfPangram(sentence string) bool {
	N := 26
	if len(sentence) < N {
		return false
	}
	set := 1<<N - 1
	for _, v := range sentence {
		mark := 1 << (v - 'a')
		if set&mark > 0 {
			set -= mark
			if set == 0 {
				return true
			}
		}
	}
	return false
}
