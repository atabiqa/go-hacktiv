package main

import "fmt"

func main() {
	cars := []string{"ford", "honda", "audi", "range rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "nissan"
	fmt.Println("cars :", cars)
	fmt.Println("newcars :", newCars)
}
