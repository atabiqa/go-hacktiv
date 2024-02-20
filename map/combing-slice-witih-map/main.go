package main

import "fmt"

func main() {
	var people = []map[string]string{
		{"name": "aira", "age": "22"},
		{"name": "nilo", "age": "27"},
		{"name": "gani", "age": "15"},
	}

	for i, person := range people {
		fmt.Printf("index: %d, name: %s, age: %s\n", i, person["name"], person["age"])
	}
}
