package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func doSomeWork() (int, error) {
	time.Sleep(10 * time.Second)
	return 0, nil
}

func main() {
	result, err := timeLimit(doSomeWork)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result)
}

func timeLimit(doSomeWork func() (int, error)) (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("work timed out")
	}
}
