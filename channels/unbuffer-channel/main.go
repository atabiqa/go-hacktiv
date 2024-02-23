package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine start sending data into channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 minute")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data", d)

	close(c1)
	time.Sleep(time.Second)
}
