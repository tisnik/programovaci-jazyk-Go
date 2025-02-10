// Základy programovacího jazyka Go
//
// - operátory pro bitové posuny kombinované
//   s přiřazením

package main

import "fmt"

func main() {
	// původní hodnota x
	var x int = 1

	// postupné násobení hodnoty uložené v x dvojkou
	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d << %2d == %4d\n", 1, i, x)
		// výpočet další hodnoty v řadě
		x <<= 1
	}

	fmt.Println()

	// původní hodnota x
	x = 10000000

	// postupné dělení hodnoty uložené v x dvojkou
	// (i musí být numerický typ bez znaménka)
	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d >> %2d == %4d\n", 1, i, x)
		// výpočet další hodnoty v řadě
		x >>= 1
	}

}
