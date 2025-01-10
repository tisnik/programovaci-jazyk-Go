// Základy programovacího jazyka Go
//
// - deklarace lokální proměnné v rámci bloku

package main

import "fmt"

func main() {
	if true {
		// proměnná, která je lokální v rámci bloku
		var counter int = 5

		// výpis původní hodnoty proměnné
		fmt.Println(counter)

		// modifikace proměnné
		counter += 1

		// výpis nové hodnoty proměnné
		fmt.Println(counter)
	}
}
