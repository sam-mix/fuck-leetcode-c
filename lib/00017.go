package lib

import "fmt"

var combinations []string
var phoneMap = make(map[string]string)

func letterCombinations(digits string) []string {

	ch := byte('a')
	for i := 2; i < 10; i++ {
		c := 3
		if i == 7 || i == 9 {
			c = 4
		}
		l := make([]byte, c)
		for j := 0; j < c; j++ {
			l[j] = ch
			ch = ch + 1
		}
		phoneMap[fmt.Sprintf("%d", i)] = string(l)
	}
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}
