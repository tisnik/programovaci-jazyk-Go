// Základy programovacího jazyka Go
//
// - deklarace globální proměnné
// - modifikace globální proměnné ve funkci

package main

import "fmt"

// proměnná, která je globální pro celý modul
var counter int = 5

func main() {
	// výpis původní hodnoty proměnné
	fmt.Println(counter)

	// modifikace proměnné
	counter += 1

	// výpis nové hodnoty proměnné
	fmt.Println(counter)
}
