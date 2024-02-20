package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":   "ata",
		"age":    "26",
		"addres": "jalan suderman",
	}
	value, exist := person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value does'nt exist")
	}

	delete(person, "age")

	value, exist = person["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("value has been deleted")
	}
}
