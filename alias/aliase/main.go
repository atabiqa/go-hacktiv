package main

import "fmt"

func main() {
	//mendeklarasikan tipe data alias bernama second
	//type nama_alias = nama_type_data

	type second = uint

	var hour second = 3600
	fmt.Printf("hour type : %T\n", hour) // =>hour type uint
}
