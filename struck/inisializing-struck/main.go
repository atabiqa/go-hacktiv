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

	var employee2 = Employee{name: "ann", age: 45, division: "human"}
	fmt.Println(employee)
	fmt.Println(employee2)
}
