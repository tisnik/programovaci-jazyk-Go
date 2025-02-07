// Základy programovacího jazyka Go
//
// - bitové posuny pro hodnoty typu int

package main

import "fmt"

func main() {
	// celočíselný typ se znaménkem
	var x int = -42

	// zobrazení původní hodnoty x
	fmt.Printf("x =     %08b\n", x)

	// zobrazení hodnoty x po jejím posunu doprava a doleva
	// vždy o jeden bit
	fmt.Printf("x>>1 =  %08b\n", x>>1)
	fmt.Printf("x<<1 =  %08b\n", x<<1)
}
