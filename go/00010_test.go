package lib

import "testing"

func TestIsMatch(t *testing.T) {
	cases := []struct {
		in struct {
			str string
			p   string
		}
		out bool
	}{
		{
			struct {
				str string
				p   string
			}{"", ".*"},
			true,
		},
	}
	for _, v := range cases {
		assertEqual(t, v.out, isMatch(v.in.str, v.in.p), v.in)
	}
}
