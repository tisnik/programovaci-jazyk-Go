// Základy programovacího jazyka Go
//
// - základní relační operátory s proměnnými
//   typu int

package main

import "fmt"

func main() {
	// dvojice proměnných typu int
	var x int = 1
	var y int = 2

	// výpis původních hodnot
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)

	fmt.Println()

	// výpis výsledků operací
	fmt.Printf("x == y: %t\n", x == y)
	fmt.Printf("x != y: %t\n", x != y)
	fmt.Printf("x > y: %t\n", x > y)
	fmt.Printf("x >= y: %t\n", x >= y)
	fmt.Printf("x < y: %t\n", x < y)
	fmt.Printf("x <= y: %t\n", x <= y)
}
