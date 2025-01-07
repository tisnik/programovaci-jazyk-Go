// Základy programovacího jazyka Go
//
// - bitové posuny o hodnotu proměnné (bez znaménka)

package main

import "fmt"

func main() {
	// původní hodnota x
	var x int = 1

	// postupné násobení hodnoty x mocninami dvojky
	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d << %2d == %4d\n", x, i, x<<i)
	}

	fmt.Println()

	// původní hodnota x
	x = 10000000

	// postupné dělení hodnoty x mocninami dvojky
	for i := uint(0); i <= 10; i++ {
		fmt.Printf("%d >> %2d == %4d\n", x, i, x>>i)
	}

}
