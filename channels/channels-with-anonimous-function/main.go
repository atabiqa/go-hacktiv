package main

import (
	"fmt"
)

func main() {
	c := make(chan string)

	students := []string{"airel", "mailo", "indah"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("student", student)
			result := fmt.Sprintf("hai, my name is %s", student)
			c <- result
		}(v)

	}

	for i := 1; i <= 3; i++ {
		print(c)
	}
	close(c)

}
