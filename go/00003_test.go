package lib

import (
	"testing"
)

func assertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}
func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{
			"abcabdc",
			4,
		},
		{
			"abcdefabdc",
			6,
		},
		{
			"ababcfgedabdc",
			7,
		},
	}
	for _, v := range cases {
		len := lengthOfLongestSubstring(v.in)
		assertEqual(t, v.out, len)
	}

}
