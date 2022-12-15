package lib

func getLucky1(s string, k int) (ans int) {
	for _, b := range s {
		b -= 'a' - 1
		ans += int(b/10 + b%10)
	}
	for i := 1; i < k && ans > 9; i++ {
		s := ans
		for ans = 0; s > 0; s /= 10 {
			ans += s % 10
		}
	}
	return
}

func getLucky(s string, k int) (sum int) {
	for _, b := range s {
		b -= 'a' - 1
		sum += int(b/10 + b%10)
	}
	for k--; k > 0 && sum > 9; k-- { // sum < 10 时可以提前退出
		s := sum
		for sum = 0; s > 0; s /= 10 {
			sum += s % 10
		}
	}
	return
}
