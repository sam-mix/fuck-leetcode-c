package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Hello world!")
		close(done)
	}()
	<-done
	fmt.Println("done")
}
