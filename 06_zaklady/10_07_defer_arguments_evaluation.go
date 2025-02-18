// Základy programovacího jazyka Go
//
// - ověření, kdy se vyhodnocují parametry předávané funkci
//   s odloženým zavoláním
// - předávání řezů

package main

import "fmt"

func function(a []int) {
	fmt.Printf("Defer %v\n", a)
}

func main() {
	// proměnná, jejíž hodnota bude předáváná do funkce
	var x = []int{1, 2, 3}
	fmt.Printf("Current x value = %v\n", x)

	// odložené zavolání funkce s předáním parametru
	defer function(x)

	// modifikace obsahu řezu
	x[0] = 0
	fmt.Printf("Current x value = %v\n", x)

	// odložené zavolání funkce s předáním parametru
	defer function(x)

	// modifikace obsahu řezu
	x[1] = 0
	fmt.Printf("Current x value = %v\n", x)

	// odložené zavolání funkce s předáním parametru
	defer function(x)

	// modifikace obsahu řezu
	x[2] = 0
	fmt.Printf("Current x value = %v\n", x)

	// odložené zavolání funkce s předáním parametru
	defer function(x)
}
