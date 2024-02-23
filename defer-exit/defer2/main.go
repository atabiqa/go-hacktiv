package main

import "fmt"

func main() {
	callDeferFunc()
	fmt.Println("hai everyone")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer function start to execute")
}
