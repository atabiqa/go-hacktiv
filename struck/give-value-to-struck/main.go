package main

import "fmt"

type Employee struct {
	name     string
	age      int
	division string
}

func main() {
	var employee Employee

	employee.name = "julie"
	employee.age = 22
	employee.division = "developer"

	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
}
