package lib

import "math"

// func reverse(x int) int {
// 	a, p := x, 1
// 	if x < 0 {
// 		a, p = -x, -1
// 	}
// 	r := 0
// 	for a > 0 {
// 		t := a % 10
// 		r = r*10 + t
// 		a /= 10
// 	}
// 	return r * p
// }

// func reverse(x int) int {
// 	a := x
// 	rbs := []byte{}
// 	if x < 0 {
// 		a = -x
// 		rbs = []byte{'-'}
// 	}
// 	flag := false
// 	for a > 0 {
// 		t := a % 10
// 		if flag || t > 0 {
// 			flag = true
// 			rbs = append(rbs, []byte(fmt.Sprintf("%d", t))...)
// 		}
// 		a /= 10
// 	}
// 	r, _ := strconv.ParseInt(string(rbs), 10, 64)
// 	return int(r)
// }

func reverse(x int) (rev int) {
	for x != 0 {
		if rev < math.MinInt32/10 || rev > math.MaxInt32/10 {
			return 0
		}
		digit := x % 10
		x /= 10
		rev = rev*10 + digit
	}
	return
}
