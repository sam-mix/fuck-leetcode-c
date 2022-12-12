package lib

import (
	"testing"
)

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
		assertEqual(t, v.out, len, v.in)
	}

}

func BenchmarkXxx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("aaaa")
	}
}
