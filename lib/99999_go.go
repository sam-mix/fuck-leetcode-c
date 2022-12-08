package lib

import (
	"fmt"
	"sort"
)

func gofor() {
	n, i := 100, 0
	for i <= n {
		n--
		i++
		fmt.Printf("%d, %d\n", n, i)
	}
	a := []int{4, 5, 2}
	sort.Ints(a)
	fmt.Println(a)
	b := "1234567"
	for i := range b {
		fmt.Println(i)
	}
	fmt.Println("done")
}
