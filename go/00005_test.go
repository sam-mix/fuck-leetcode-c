package lib

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{
			in:  "abab",
			out: "aba",
		},
		{
			in:  "cbbd",
			out: "bb",
		},
		{
			in:  "bbbb",
			out: "bbbb",
		},
		{
			in:  "bbbbbb",
			out: "bbbbbb",
		},
		{
			in:  "abccba",
			out: "abccba",
		},
		{
			in:  "bb",
			out: "bb",
		},
		{
			in:  "ccc",
			out: "ccc",
		},
	}
	for _, v := range cases {
		assertEqual(t, v.out, longestPalindrome(v.in), v.in)
	}
}

func TestLongestPalindromeOn(t *testing.T) {
	cases := []struct {
		in  string
		out string
	}{
		{
			in:  "abab",
			out: "bab",
		},
		{
			in:  "cbbd",
			out: "bb",
		},
		{
			in:  "bbbb",
			out: "bbbb",
		},
		{
			in:  "bbbbbb",
			out: "bbbbbb",
		},
		{
			in:  "abccba",
			out: "abccba",
		},
		{
			in:  "bb",
			out: "bb",
		},
		{
			in:  "ccc",
			out: "ccc",
		},
	}
	for _, v := range cases {
		assertEqual(t, v.out, longestPalindromeOn(v.in), v.in)
	}
}
