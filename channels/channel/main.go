package main

import "fmt"

func main() {
	c := make(chan string)

	go introduce("airel", c)

	go introduce("fanani", c)

	go introduce("geodani", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("hai, my name is %s", student)

	c <- result
}
