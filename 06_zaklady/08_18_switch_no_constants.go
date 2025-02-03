package main

import "fmt"

func classify(x int, zero_value int) string {
	switch x {
	case zero_value:
		return "nula"
	case 2, 4, 6, 8:
		return "sudé číslo"
	case 1, 3, 5, 7, 9:
		return "liché číslo"
	default:
		return "mimo zkoumaný rozsah"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Println(x, classify(x, 0))
	}
}
