// Základy programovacího jazyka Go
//
// - základní aritmetické operátory použité s proměnnými
//   typu int

package main

import "fmt"

func main() {
	// dvojice proměnných typu int
	var x int = 10
	var y int = 3

	// výpis původních hodnot proměnných
	fmt.Printf("x = %d\n", x)
	fmt.Printf("y = %d\n", y)

	fmt.Println()

	// výpis výsledků aritmetických operací s proměnnými
	fmt.Printf("x+y = %d\n", x+y)
	fmt.Printf("x-y = %d\n", x-y)
	fmt.Printf("x*y = %d\n", x*y)
	fmt.Printf("x/y = %d\n", x/y)
	fmt.Printf("x%%y = %d\n", x%y)
}
