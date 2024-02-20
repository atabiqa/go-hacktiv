package main

import "fmt"

//cara pertama menambhakan isi slice make function

// func main() {
// 	var fruits = make([]string, 3)
// 	_ = fruits

// 	fruits[0]="apple"
// 	fruits[1]="banana"
// 	fruits[3]="mango"
// 	fmt.Printf("%#v", fruits)
// }

//cara kedua dengan appand
func main() {
	var fruits = make([]string, 3)
	fruits = append(fruits, "apple", "banana", "mango")
	fmt.Printf("%#v", fruits)
}
