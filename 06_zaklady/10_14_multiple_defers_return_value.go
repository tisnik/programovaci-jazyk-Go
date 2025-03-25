// Základy programovacího jazyka Go
//
// - manipulace s návratovými hodnotami přímo v konstrukci defer

package main

import "fmt"

func function(x int) (i int) {
	// odložené nastavení návratové hodnoty
	defer func() { i += 1 }()
	defer func() { i *= 10 }()
	// explicitní návrat s hodnotou 1
	return x
}

func main() {
	fmt.Printf("Return value of function: %d\n", function(100))
}
