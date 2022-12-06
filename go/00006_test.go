package lib

import "testing"

func TestConvert(t *testing.T) {
	cases := []struct {
		in struct {
			str string
			col int
		}
		out string
	}{
		{
			in: struct {
				str string
				col int
			}{"abcd", 4},
			out: "abcd",
		},
		{
			in: struct {
				str string
				col int
			}{"abcdefg", 3},
			// a    e
			// b  d    f
			// c .       g
			out: "aebdfcg",
		},
		{
			in: struct {
				str string
				col int
			}{"abcd", 2},
			out: "acbd",
		},
		{
			in: struct {
				str string
				col int
			}{"abcd", 1},
			out: "abcd",
		},
	}
	for _, v := range cases {
		assertEqual(t, v.out, convert(v.in.str, v.in.col), v.in)
	}
}
