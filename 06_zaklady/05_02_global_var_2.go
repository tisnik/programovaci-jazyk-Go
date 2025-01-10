// Základy programovacího jazyka Go
//
// - deklarace globální proměnné
// - proměnná má výchozí (nulovou) hodnotu
// - modifikace globální proměnné ve funkci

package main

import "fmt"

// proměnná, která je globální pro celý modul
// proměnná má výchozí (nulovou) hodnotu
var counter int

func main() {
	// výpis původní hodnoty proměnné
	fmt.Println(counter)

	// modifikace proměnné
	counter += 1

	// výpis nové hodnoty proměnné
	fmt.Println(counter)
}
