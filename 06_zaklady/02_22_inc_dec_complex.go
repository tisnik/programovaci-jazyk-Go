// Základy programovacího jazyka Go
//
// - operátory ++ a -- použité pro proměnné typu complex

package main

import "fmt"

func main() {
	// původní hodnota z
	var z complex64 = 1 + 2i

	// výpis původní hodnoty
	fmt.Printf("z = %v\n", z)

	// zvýšení hodnoty uložené v proměnné x o jedničku
	z++

	// výpis nové hodnoty
	fmt.Printf("z = %v\n", z)

	// snížení hodnoty uložené v proměnné x o jedničku
	z--

	// výpis nové hodnoty
	fmt.Printf("z = %v\n", z)
}
