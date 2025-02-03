// Základy programovacího jazyka Go
//
// - programová smyčka for s podmínkou na začátku

package main

import "fmt"

func main() {
	i := 10

	// podmínka vyhodnocovaná před každou iterací
	for i != 0 {
		fmt.Println("Counting:", i)
		// změna počitadla smyčky
		i--
	}
}
