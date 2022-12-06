package lib

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in  int
		out int
	}{
		{
			in:  123456789,
			out: 0,
		},
		{
			in:  -100,
			out: -1,
		},
		{
			in:  -100000000,
			out: -1,
		},
		{
			in:  901000,
			out: 109,
		},
		{
			in:  1534236469,
			out: 0,
		},
	}
	for _, v := range cases {
		assertEqual(t, v.out, reverse(v.in), v.in)
	}
}
