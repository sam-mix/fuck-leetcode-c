package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer func() {
			if Any := recover(); Any != nil {
				fmt.Println("Any: ", Any)
			}
		}()
		panic("hello")
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("hello world")
}
