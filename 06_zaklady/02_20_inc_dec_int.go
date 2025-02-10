// Základy programovacího jazyka Go
//
// - operátory ++ a -- použité pro proměnné typu int

package main

import "fmt"

func main() {
	// původní hodnota x
	var x int = 1

	// výpis původní hodnoty
	fmt.Printf("x = %d\n", x)

	// zvýšení hodnoty uložené v proměnné x o jedničku
	x++

	// výpis nové hodnoty
	fmt.Printf("x = %d\n", x)

	// snížení hodnoty uložené v proměnné x o jedničku
	x--

	// výpis nové hodnoty
	fmt.Printf("x = %d\n", x)
}
