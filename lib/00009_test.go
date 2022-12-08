package lib

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in  int
		out bool
	}{
		{123321, true},
		{12321, true},
		{123, false},
		{100, false},
		{-1, false},
		{0, true},
	}

	for _, v := range cases {
		assertEqual(t, v.out, isPalindrome(v.in), v.in)
	}
}
