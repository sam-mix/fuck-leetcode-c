package lib

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	cases := []struct {
		in  []string
		out string
	}{
		{[]string{"", "", ""}, ""},
		{[]string{"flower", "flow", "flight"}, "fl"},
	}
	for _, v := range cases {
		assertEqual(t, v.out, longestCommonPrefix(v.in), v.in)
	}
}
