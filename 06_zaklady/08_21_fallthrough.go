// Základy programovacího jazyka Go
//
// - řídicí konstrukce switch a fallthrough

package main

import "fmt"

func sayNumber(x int) {
	switch x {
	case 0:
		fmt.Println("nula")
		fallthrough
	case 1:
		fmt.Println("jedna")
		fallthrough
	case 2:
		fmt.Println("dva")
		fallthrough
	case 3:
		fmt.Println("tři")
		fallthrough
	case 4:
		fmt.Println("čtyři")
		fallthrough
	case 5:
		fmt.Println("pět")
		fallthrough
	case 6:
		fmt.Println("šest")
		fallthrough
	case 7:
		fmt.Println("sedm")
		fallthrough
	case 8:
		fmt.Println("osm")
		fallthrough
	case 9:
		fmt.Println("devět")
		fallthrough
	case 10:
		fmt.Println("deset")
		fallthrough
	default:
		fmt.Println("mimo rozsah")
	}
}

func main() {
	for x := 5; x <= 10; x++ {
		sayNumber(x)
		fmt.Println()
	}
}
