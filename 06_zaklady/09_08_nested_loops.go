// Základy programovacího jazyka Go
//
// - vnořené programové smyčky
// - zobrazení tabulky malé násobilky

package main

import "fmt"

func main() {
	// vnější počítaná smyčka pro hodnoty i od 0 do 9
	for i := range 10 {
		// vnitřní počítaná smyčka pro hodnoty i od 0 do 9
		for j := range 10 {
			fmt.Printf("%3d ", (i+1)*(j+1))
		}
		// odřádkování
		fmt.Println()
	}
}
