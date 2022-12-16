package lib

import "strings"

func isPalindromeMe(s string) bool {
	lower := func(a byte) (bool, byte) {
		if (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
			return true, a
		}
		if a >= 'A' && a <= 'Z' {
			return true, a - 26
		}

		return false, ' '
	}
	for i, j := 0, len(s)-1; i < j; {
		ok, a := lower(s[i])
		if !ok {
			i++
			continue
		}
		ok, b := lower(s[j])
		if !ok {
			j--
			continue
		}
		if a != b {
			return false
		}
		i++
		j--
	}
	return true
}

func isPalindromeMe1(s string) bool {
	s = strings.ToLower(s)
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isalnum(s[left]) {
			left++
		}
		for left < right && !isalnum(s[right]) {
			right--
		}
		if left < right {
			if s[left] != s[right] {
				return false
			}
			left++
			right--
		}
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
