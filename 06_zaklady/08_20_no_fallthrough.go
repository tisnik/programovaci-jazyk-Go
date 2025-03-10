// Základy programovacího jazyka Go
//
// - řídicí konstrukce switch a implicitní příkazy break

package main

import "fmt"

func sayNumber(x int) {
	switch x {
	case 0:
		fmt.Println("nula")
	case 1:
		fmt.Println("jedna")
	case 2:
		fmt.Println("dva")
	case 3:
		fmt.Println("tři")
	case 4:
		fmt.Println("čtyři")
	case 5:
		fmt.Println("pět")
	case 6:
		fmt.Println("šest")
	case 7:
		fmt.Println("sedm")
	case 8:
		fmt.Println("osm")
	case 9:
		fmt.Println("devět")
	case 10:
		fmt.Println("deset")
	default:
		fmt.Println("mimo rozsah")
	}
}

func main() {
	for x := 0; x <= 10; x++ {
		sayNumber(x)
	}
}
