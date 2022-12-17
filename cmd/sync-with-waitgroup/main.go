package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		doThing1()
	}()
	go func() {
		defer wg.Done()
		doThing2()
	}()
	go func() {
		defer wg.Done()
		doThing3()
	}()
	wg.Wait()
}

func doThing1() {
	fmt.Println("doThing1")
}

func doThing2() {
	fmt.Println("doThing2")
}

func doThing3() {
	fmt.Println("doThing3")
}
