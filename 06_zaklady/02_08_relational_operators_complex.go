// Základy programovacího jazyka Go
//
// - základní relační operátory s proměnnými
//   typu complex64

package main

import "fmt"

func main() {
	// dvojice proměnných typu complex64
	var x complex64 = 1 + 0i
	var y complex64 = 1 + 1i

	// výpis původních hodnot
	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)

	fmt.Println()

	// výpis výsledků operací
	fmt.Printf("x == y: %t\n", x == y)
	fmt.Printf("x != y: %t\n", x != y)
}
