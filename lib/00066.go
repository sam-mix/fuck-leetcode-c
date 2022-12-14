package lib

func plusOne(digits []int) []int {
	one := 1 /*进位*/
	for i := len(digits) - 1; i > -1; i-- {
		temp := digits[i] + one
		if temp > 9 {
			digits[i] = temp - 10
			one = 1
		} else {
			digits[i] = temp
			one = 0
			break
		}
	}
	if one == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
