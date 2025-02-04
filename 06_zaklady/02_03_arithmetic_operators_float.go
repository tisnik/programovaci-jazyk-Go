// Základy programovacího jazyka Go
//
// - základní aritmetické operátory s proměnnými
//   typu float32

package main

import "fmt"

func main() {
	// dvojice proměnných typu float32
	var x float32 = 1.1
	var y float32 = 2.2

	// výpis původních hodnot proměnných
	fmt.Printf("x = %f\n", x)
	fmt.Printf("y = %f\n", y)

	fmt.Println()

	// výpis výsledků aritmetických operací s proměnnými
	fmt.Printf("x+y = %f\n", x+y)
	fmt.Printf("x-y = %f\n", x-y)
	fmt.Printf("x*y = %f\n", x*y)
	fmt.Printf("x/y = %f\n", x/y)
}
