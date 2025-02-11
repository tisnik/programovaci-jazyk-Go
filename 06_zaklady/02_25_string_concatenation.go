// Základy programovacího jazyka Go
//
// - spojení řetězců operátorem +

package main

import "fmt"

func main() {
	// dvojice proměnných typu string
	var x string = "Hello "
	var y string = "world!"

	// výpis původních hodnot proměnných
	fmt.Printf("x = '%s'\n", x)
	fmt.Printf("y = '%s'\n", y)

	fmt.Println()

	// spojení řetězců operátorem +
	fmt.Printf("x+y = %s\n", x+y)
}
