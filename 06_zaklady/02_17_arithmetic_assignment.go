// Základy programovacího jazyka Go
//
// - základní aritmetické operátory kombinované
//   s přiřazením

package main

import "fmt"

func main() {
	// původní hodnota x
	var x int = 1

	// zobrazení původní hodnoty uložené v x
	fmt.Printf("x = %d\n", x)

	// modifikace x aritmetickou operací +
	x += 10
	fmt.Printf("x = %d\n", x)

	// modifikace x aritmetickou operací *
	x *= 2
	fmt.Printf("x = %d\n", x)

	// modifikace x aritmetickou operací /
	x /= 3
	fmt.Printf("x = %d\n", x)

	// modifikace x aritmetickou operací %
	x %= 2
	fmt.Printf("x = %d\n", x)
}
