package main

func main() {
	var i chan int = nil
	go func() {
		<-i
	}()
	i <- 10
	// time.Sleep(10 * time.Second)
}
