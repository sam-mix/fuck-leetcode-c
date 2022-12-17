package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func(v int) {
			time.Sleep(time.Duration(v) * time.Millisecond)
			ch <- v
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
