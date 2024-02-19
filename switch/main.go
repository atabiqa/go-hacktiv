package main

import "fmt"

func main() {
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesom")
	default:
		fmt.Println("not bad")
	}
}
