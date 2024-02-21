package main

import (
	"fmt"
	"strings"
)

func main() {
	profil("airel", "pasta", "ayam geprek", "ikan roa", "sate padang")
}

func profil(name string, favFood ...string) {
	mergeFavFoods := strings.Join(favFood, ",")

	fmt.Println("Hello there!!!I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
