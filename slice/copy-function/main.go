package main

import "fmt"

func main() {
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pinneple", "starfruits"}
	nn := copy(fruits1, fruits2)
	fmt.Println("fruits1 =>", fruits1)
	fmt.Println("fruits2 =>", fruits2)
	fmt.Println("copied element =>", nn)
}
