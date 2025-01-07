// Základy programovacího jazyka Go
//
// - základní relační operátory s proměnnými
//   typu řetězec

package main

import "fmt"

func main() {
	// dvojice proměnných typu string
	var x string = "foo"
	var y string = "bar"

	// výpis původních hodnot
	fmt.Printf("x = '%s'\n", x)
	fmt.Printf("y = '%s'\n", y)

	fmt.Println()

	// výpis výsledků operací
	fmt.Printf("x == y: %t\n", x == y)
	fmt.Printf("x != y: %t\n", x != y)
	fmt.Printf("x > y: %t\n", x > y)
	fmt.Printf("x >= y: %t\n", x >= y)
	fmt.Printf("x < y: %t\n", x < y)
	fmt.Printf("x <= y: %t\n", x <= y)
}
