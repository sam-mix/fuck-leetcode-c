// package main

// import "sam-mix.ltd/fuck-leetcode-c/lib"

// func main() {
// 	lib.IntToRoman(123423451234)
// }

package main

import "fmt"

func main() {
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)
	fmt.Println(s)
	s = s[0 : len(s)-1]
	fmt.Println(s)
	fmt.Println(len(s))
}
