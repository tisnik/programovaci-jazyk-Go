// Základy programovacího jazyka Go
//
// - bitové posuny pro hodnoty typu byte

package main

import "fmt"

func main() {
	// celočíselný typ bez znaménka
	var x byte = 42

	// zobrazení původní hodnoty x
	fmt.Printf("x =     %08b\n", x)

	// zobrazení hodnoty x po jejím posunu doprava a doleva
	// vždy o jeden bit
	fmt.Printf("x>>1 =  %08b\n", x>>1)
	fmt.Printf("x<<1 =  %08b\n", x<<1)
}
