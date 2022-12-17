package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	result := channelProcess(ch)
	fmt.Println(result)

}

func channelProcess(ch chan int) []int {
	const conc = 10
	results := make(chan int, conc)
	for i := 0; i < conc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	var out []int
	for i := 0; i < conc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
	return 2 * v
}
