// Základy programovacího jazyka Go
//
// - řídicí konstrukce switch - simulace chování C.

package main

import "fmt"

func classify(x int) string {
	switch x {
	case 0:
		return "nula"
	case 2:
		fallthrough
	case 4:
		fallthrough
	case 6:
		fallthrough
	case 8:
		return "sudé číslo"
	case 1:
		fallthrough
	case 3:
		fallthrough
	case 5:
		fallthrough
	case 7:
		fallthrough
	case 9:
		return "liché číslo"
	default:
		return "mimo zkoumaný rozsah"
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Println(x, classify(x))
	}
}
