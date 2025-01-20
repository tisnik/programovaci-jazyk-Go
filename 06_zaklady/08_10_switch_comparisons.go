// Základy programovacího jazyka Go
//
// - konstrukce switch s podmínkami

package main

import "fmt"

func main() {
	i := 5
	switch {
	case i < 5:
		fmt.Println(i, "< 5")
	case i > 5:
		fmt.Println(i, "> 5")
	default:
		fmt.Println(i, "ani větší než 5, ani menší než 5")
	}
}
