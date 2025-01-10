// Základy programovacího jazyka Go
//
// - deklarace dvou globálních proměnných na jediném řádku

package main

import "fmt"

// deklarace dvou globálních proměnných na jediném řádku
var X, Y float32 = 0.0, 1.0

func main() {
	// výpis obou globálních proměnných
	fmt.Println(X, Y)
}
