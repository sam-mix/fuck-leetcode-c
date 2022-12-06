package lib

import "testing"

func TestMyAtoi(t *testing.T) {
	cases := []struct {
		in  string
		out int
	}{
		{"-1 .   dasfewasd", -1},
		{"words and 987", 0},
		{".......1 .   dasfewasd", 0},
		{"xxxxxx1 .   dasfewasd", 0},
		{"zzzz 1 .   dasfewasd", 0},
		{"xdvjajga 1 .   dasfewasd", 0},
		{" . 1.. .   dasfewasd", 0},
	}
	for _, v := range cases {
		assertEqual(t, v.out, myAtoi(v.in), v.in)
	}
}
