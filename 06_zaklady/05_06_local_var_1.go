// Základy programovacího jazyka Go
//
// - deklarace lokální proměnné v rámci funkce

package main

import "fmt"

func main() {
	// proměnná, která je lokální v rámci funkce
	var counter int = 5

	// výpis původní hodnoty proměnné
	fmt.Println(counter)

	// modifikace proměnné
	counter += 1

	// výpis nové hodnoty proměnné
	fmt.Println(counter)
}
