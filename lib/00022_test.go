package lib

import "testing"

func TestGenerateParenthesis1(t *testing.T) {
	cases := []struct {
		in  int
		out []string
	}{
		{0, []string{}},
		{1, []string{"()"}},
		{2, []string{"(())",
			"()()"}},
		{3, []string{"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()"}},
		{4, []string{"(((())))",
			"((()()))",
			"((())())",
			"((()))()",
			"(()(()))",
			"(()()())",
			"(()())()",
			"(())(())",
			"(())()()",
			"()((()))",
			"()(()())",
			"()(())()",
			"()()(())",
			"()()()()"}},
	}

	for _, v := range cases {
		// assertEqual(t, v.out, generateParenthesis(v.in), v.in)
		rs, temes := generateParenthesis1(v.in)
		for _, r := range rs {
			t.Logf("%s\n", r)
		}
		t.Log("tests", temes, "\n")
	}
}
