// Základy programovacího jazyka Go
//
// - základní aritmetické operátory s proměnnými
//   typu complex

package main

import "fmt"

func main() {
	// dvojice proměnných typu complex64
	var x complex64 = 1 + 1i
	var y complex64 = 2 + 2i

	// výpis původních hodnot
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)

	fmt.Println()

	// výpis výsledků operací
	fmt.Printf("x+y = %v\n", x+y)
	fmt.Printf("x-y = %v\n", x-y)
	fmt.Printf("x*y = %v\n", x*y)
	fmt.Printf("x/y = %v\n", x/y)
}
