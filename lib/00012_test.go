package lib

import "testing"

func TestIntToRoman(t *testing.T) {
	cases := []struct {
		in  int
		out string
	}{
		{3, "III"},
	}
	for _, v := range cases {
		assertEqual(t, v.out, IntToRoman(v.in), v.in)

	}
}
