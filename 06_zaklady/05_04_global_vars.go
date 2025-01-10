// Základy programovacího jazyka Go
//
// - deklarace dvou globálních proměnných v bloku var

package main

import "fmt"

// deklarace dvou globálních proměnných v bloku var
var (
	X float32 = 0.0
	Y float32 = 1.0
)

func main() {
	// výpis obou globálních proměnných
	fmt.Println(X, Y)
}
