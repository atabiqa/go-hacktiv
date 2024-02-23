package main

import "fmt"

func main() {
	c := make(chan string, 3)
	fmt.Println("c", c)
}
