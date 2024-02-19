package main

import "fmt"

func main() {
	var currentYear = 2021

	if age := currentYear - 1998; age < 17 {
		fmt.Println("kamu belum boleh membuat kartau sim")
	} else {
		fmt.Println("boleh membuat kartu sim")
	}
}
