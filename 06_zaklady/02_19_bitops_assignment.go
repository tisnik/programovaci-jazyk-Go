// Základy programovacího jazyka Go
//
// - operátory pro bitové operace kombinované
//   s přiřazením (použito u proměnné typu byte)

package main

import "fmt"

func main() {
	// původní hodnota x (osmibitové celé číslo bez znaménka)
	var x byte = 0b11110000
	fmt.Printf("original x:   %08b\n", x)

	// nastavení bitu číslo 0 na jedničku
	x |= 0b00000001
	fmt.Printf("set bit #0:   %08b\n", x)

	// vynulování bitu číslo 7
	x &= 0b01111111
	fmt.Printf("reset bit #7: %08b\n", x)

	// maskování hodnoty operací AND NOT
	// (maska je negována a poté se provede bitová operace AND)
	x &^= 0b11110000
	fmt.Printf("masked:       %08b\n", x)
}
