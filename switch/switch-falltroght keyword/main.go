package main

import "fmt"

func main() {
	var score = 6

	switch {
	case score == 8:
		fmt.Println("perfect")
	case (score < 8) && (score < 3):
		fmt.Println("not bad")
		fallthrough
	case score < 5:
		fmt.Println("it is ok, but please studei harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you done have a good score yet")
		}
	}
}
