// Základy programovacího jazyka Go
//
// - řídicí konstrukce switch - pracuje jinak než v C!

package main

import "fmt"

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2:
	case 4:
	case 6:
	case 8:
		return "sudé číslo"
	case 1:
	case 3:
	case 5:
	case 7:
	case 9:
		return "liché číslo"
	default:
		return "mimo zkoumaný rozsah"
	}
	return "X"
}

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Println(x, classify(x))
	}
}
