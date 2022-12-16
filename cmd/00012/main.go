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
	mx(s)
	fmt.Println(s)
	ss := "Aabc"
	ss = "s" + ss
	fmt.Println(ss)
	for _, v := range ss {
		fmt.Println(v)
	}
	fmt.Println("")
	fmt.Println('a')
	fmt.Println('A')

	fmt.Println("")
	fmt.Println('0')

	a := append([]int{}, 1)
	fmt.Println(a)
}

func mx(l []int) {
	l[0] = 5
}
