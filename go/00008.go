package lib

import "math"

// func myAtoi(s string) int {
// 	a, p := 0, 1
// 	for _, v := range s {
// 		if a == 0 {
// 			if v < '0' && v > '9' && v != ' ' && v != '-' {
// 				return 0
// 			}
// 			if v == '-' {
// 				p = -1
// 			}
// 		}
// 		if v >= '0' && v <= '9' {
// 			a = a*10 + int(v-'0')
// 			if p*a < -2<<32 {
// 				return -2 << 32
// 			}
// 			if p*a > 2<<32-1 {
// 				return 2<<32 - 1
// 			}
// 		}
// 	}

// 	return a * p
// }

func myAtoi(s string) int {
	ans := 0
	negative := false
	index := 0
	for index < len(s) && s[index] == ' ' {
		index++
	}
	if index >= len(s) {
		return 0
	}

	if s[index] == '+' || s[index] == '-' {
		if s[index] == '-' {
			negative = true
		}
		index++
	}
	for ; index < len(s); index++ {
		num := s[index] - '0'
		if num < 0 || num > 9 {
			break
		}
		if ans > math.MaxInt32/10 || ans == math.MaxInt32/10 && num > 7 {
			if negative {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		}
		ans = ans*10 + int(num)
	}
	if negative {
		ans = -ans
	}
	return ans
}
