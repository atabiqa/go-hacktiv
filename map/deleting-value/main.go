package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":   "ata",
		"age":    "26",
		"addres": "jalan suderman",
	}

	fmt.Println("before deleting:", person)
	delete(person, "age")
	fmt.Println("after deleting:", person)
}
