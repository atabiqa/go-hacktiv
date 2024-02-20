package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":   "ata",
		"age":    "12",
		"addres": "magelang",
	}

	fmt.Println("name:", person["name"])
	fmt.Println("age:", person["age"])
	fmt.Println("addres", person["addres"])
}
