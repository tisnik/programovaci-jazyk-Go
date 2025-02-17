// Základy programovacího jazyka Go
//
// - anonymní funkce s odloženým voláním
// - zjednodušený zápis této funkce

package main

import "fmt"

func main() {
	// anonymní funkce s odloženým voláním
	defer func() { fmt.Println("Finished") }()

	// simulace nějaké činnosti
	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
