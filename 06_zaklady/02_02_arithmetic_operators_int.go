// Základy programovacího jazyka Go
//
// - základní aritmetické operátory s proměnnými
//   typu int

package main

import "fmt"

func main() {
	// dvojice proměnných typu int
	var x int = 10
	var y int = 3

	// výpis původních hodnot
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)

	fmt.Println()

	// výpis výsledků operací
	fmt.Printf("x+y = %d\n", x+y)
	fmt.Printf("x-y = %d\n", x-y)
	fmt.Printf("x*y = %d\n", x*y)
	fmt.Printf("x/y = %d\n", x/y)
	fmt.Printf("x%%y = %d\n", x%y)
}
