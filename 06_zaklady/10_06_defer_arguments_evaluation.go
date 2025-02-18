// Základy programovacího jazyka Go
//
// - ověření, kdy se vyhodnocují parametry předávané funkci
//   s odloženým zavoláním

package main

import "fmt"

func function(i int) {
	fmt.Printf("Defer %2d\n", i)
}

func main() {
	// proměnná, jejíž hodnota bude předáváná do funkce
	x := 0
	fmt.Printf("Current x value = %2d\n", x)

	// odložené zavolání funkce s předáním parametru
	defer function(x)

	// modifikace obsahu proměnné
	x++
	fmt.Printf("Current x value = %2d\n", x)

	// odložené zavolání funkce s předáním parametru
	defer function(x)

	// modifikace obsahu proměnné
	x++
	fmt.Printf("Current x value = %2d\n", x)

	// odložené zavolání funkce s předáním parametru
	defer function(x)
}
