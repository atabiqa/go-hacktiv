package main

import "fmt"

func main() {
	var person = map[string]string{
		"name":   "ata",
		"age":    "26",
		"addres": "jalan suderman",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
