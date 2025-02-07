// Základy programovacího jazyka Go
//
// - logické operátory (pro hodnoty typu bool)

package main

import "fmt"

func main() {
	// dvojice proměnných typu bool(ean)
	var x bool = true
	var y bool = false

	// výpis původních hodnot
	fmt.Printf("x = %t\n", x)
	fmt.Printf("y = %t\n", y)

	fmt.Println()

	// výpis výsledků operací
	fmt.Printf("not x = %t\n", !x)
	fmt.Printf("not y = %t\n", !y)
	fmt.Printf("x && y = %t\n", x && y)
	fmt.Printf("x || y = %t\n", x || y)
}
