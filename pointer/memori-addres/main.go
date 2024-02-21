package main

import "fmt"

func main() {
	var firstNumber int = 4
	var secondNumber *int = &firstNumber

	fmt.Println("first number (value):", firstNumber)
	fmt.Println("first number (memori addres):", &firstNumber)

	fmt.Println("second number (value):", *secondNumber)
	fmt.Println("second number (memori addres):", secondNumber)
}
