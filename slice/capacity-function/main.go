package main

import (
	"fmt"
	"strings"
)

func main() {
	var fruits1 = []string{"apple", "mango", "durian", "banana"}
	fmt.Println("fruits1 cap:", cap(fruits1)) //4
	fmt.Println("fruits1 cap:", len(fruits1)) //4

	fmt.Println(strings.Repeat("#", 20))

	var fruits2 = fruits1[0:3]
	fmt.Println("fruits1 cap:", cap(fruits2)) //4
	fmt.Println("fruits1 cap:", len(fruits2)) //3

	fmt.Println(strings.Repeat("#", 20))

	var fruits3 = fruits1[01:]
	fmt.Println("fruits1 cap:", cap(fruits3)) //3
	fmt.Println("fruits1 cap:", len(fruits3)) //3
}
