// Základy programovacího jazyka Go
//
// - základní relační operátory s proměnnými
//   typu boolean

package main

import "fmt"

func main() {
	// dvojice proměnných typu bool(ean)
	var x bool = true
	var y bool = false

	// výpis původních hodnot proměnných
	fmt.Printf("x = %t\n", x)
	fmt.Printf("y = %t\n", y)

	fmt.Println()

	// výpis výsledků relačních operací s proměnnými
	fmt.Printf("x == y: %t\n", x == y)
	fmt.Printf("x != y: %t\n", x != y)
	fmt.Printf("x > y: %t\n", x > y)
	fmt.Printf("x >= y: %t\n", x >= y)
	fmt.Printf("x < y: %t\n", x < y)
	fmt.Printf("x <= y: %t\n", x <= y)
}
