package lib

import "fmt"

func gofor() {
	n, i := 100, 0
	for i <= n {
		n--
		i++
		fmt.Printf("%d, %d\n", n, i)
	}
	fmt.Println("done")
}
