// Základy programovacího jazyka Go
//
// - operátory ++ a -- použité pro proměnné typu float32

package main

import "fmt"

func main() {
	// původní hodnota x
	var x float32 = 3.14

	// výpis původní hodnoty
	fmt.Printf("x = %f\n", x)

	// zvýšení hodnoty uložené v proměnné x o jedničku
	x++

	// výpis nové hodnoty
	fmt.Printf("x = %f\n", x)

	// snížení hodnoty uložené v proměnné x o jedničku
	x--

	// výpis nové hodnoty
	fmt.Printf("x = %f\n", x)
}
