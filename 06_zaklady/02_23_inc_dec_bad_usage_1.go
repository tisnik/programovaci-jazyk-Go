// Základy programovacího jazyka Go
//
// - nekorektní způsob použití operací ++ a --

package main

import "fmt"

func main() {
	// původní hodnota x
	var x int = 1
	fmt.Printf("x = %d\n", x)

	// pokus o prefixový zápis operace ++
	++x
	fmt.Printf("x = %d\n", x)

	// pokus o prefixový zápis operace --
	--x
	fmt.Printf("x = %d\n", x)
}
