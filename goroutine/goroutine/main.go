package main

import (
	"fmt"
	"runtime"
)

func main() {
	go firstProcess(8)
	secondProcess(8)

	fmt.Println("No, of Goroutines:", runtime.NumGoroutine())
	fmt.Println("main excecution ended")
}

func firstProcess(index int) {
	fmt.Println("First process func started")
	for i := 1; i <= index; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("First process func ended")
}

func secondProcess(index int) {
	fmt.Println("Second process func started")
	for j := 1; j <= index; j++ {
		fmt.Println("j=", j)
	}
	fmt.Println("Second process func ended")
}
