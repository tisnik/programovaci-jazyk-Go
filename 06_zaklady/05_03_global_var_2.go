// Základy programovacího jazyka Go
//
// - deklarace globální proměnné
// - proměnná má výchozí (nulovou) hodnotu
// - proměnná je viditelná mimo rámec balíčku

package main

import "fmt"

// proměnná, která je globální pro celý balíček
// proměnná má výchozí (nulovou) hodnotu
// a je viditelná i mimo rámec balíčku
var Counter int

func main() {
	// výpis původní hodnoty proměnné
	fmt.Println(Counter)

	// modifikace proměnné
	Counter += 1

	// výpis nové hodnoty proměnné
	fmt.Println(Counter)
}
