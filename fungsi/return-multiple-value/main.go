package main

import (
	"fmt"
	"math"
)

func main() {
	var diameter float64 = 15

	var area, circumference = calculated(diameter)

	fmt.Println("area :", area)
	fmt.Println("circumference:", circumference)
}

func calculated(d float64) (float64, float64) {
	// menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	//menghitung keliling
	var circumference = math.Pi * d

	return area, circumference
}
