// Základy programovacího jazyka Go
//
// - operátory ++ a -- použité pro proměnné typu float32

package main

import "fmt"

func main() {
	// původní hodnota x
	var x float32 = 3.14

	// výpis původní hodnoty
	fmt.Printf("y = %f\n", x)

	// zvýšení hodnoty x o jedničku
	x++

	// výpis nové hodnoty
	fmt.Printf("y = %f\n", x)
}
