package lib

import "testing"

func TestGetLucky(t *testing.T) {
	cases := []struct {
		in struct {
			s string
			k int
		}
		out int
	}{
		{struct {
			s string
			k int
		}{"a", 1}, 1},
		{struct {
			s string
			k int
		}{"ab", 1}, 3},
		{struct {
			s string
			k int
		}{"ab", 0}, 12},
	}
	for _, v := range cases {
		assertEqual(t, v.out, getLucky(v.in.s, v.in.k), v.in)
	}
}
