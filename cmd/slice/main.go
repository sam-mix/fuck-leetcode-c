package main

import "fmt"

func main() {
	var a []int
	oldcap := cap(a)
	fmt.Println(len(a), cap(a))
	for i := 0; i < 9999999999; i++ {
		a = append(a, i)
		if cap(a) != oldcap {
			fmt.Println(len(a), cap(a))
			oldcap = cap(a)
		}
	}

}
