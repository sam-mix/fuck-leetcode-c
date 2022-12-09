package lib

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	res := []string{}
	var dfs func(int, int, string)
	dfs = func(lRemain int, rRemain int, path string) {
		if 2*n == len(path) {
			res = append(res, path)
			return
		}
		if lRemain > 0 {
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			dfs(lRemain, rRemain-1, path+")")
		}
	}
	dfs(n, n, "")
	return res
}

func generateParenthesis1(n int) ([]string, int) {
	res := []string{}
	var dfs func(int, int, string)
	i := 0
	dfs = func(lRemain int, rRemain int, path string) {
		fmt.Println("path: ", path)
		i++
		if 2*n == len(path) {
			res = append(res, path)
			return
		}
		if lRemain > 0 {
			fmt.Println("lRemain > 0 :", i)
			dfs(lRemain-1, rRemain, path+"(")
		}
		if lRemain < rRemain {
			fmt.Println("lRemain < rRemain :", i)
			dfs(lRemain, rRemain-1, path+")")
		}
	}

	dfs(n, n, "")

	return res, i
}
