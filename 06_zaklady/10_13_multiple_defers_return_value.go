// Základy programovacího jazyka Go
//
// - manipulace s návratovýi hodnotami přímo v konstrukci defer

package main

import "fmt"

func function() (i int) {
	// odložené nastavení návratové hodnoty
	defer func() { i *= 10 }()
	defer func() { i = 2 }()
	// explicitní návrat s hodnotou 1
	return 1
}

func main() {
	fmt.Printf("Return value of function: %d\n", function())
}
