// Základy programovacího jazyka Go
//
// - vnořené programové smyčky
// - zobrazení tabulky malé násobilky

package main

import "fmt"

func main() {
	// vnější počítaná smyčka pro hodnoty i od 0 do 9
	for i := 1; i <= 10; i++ {
		// vnitřní počítaná smyčka pro hodnoty i od 0 do 9
		for j := 1; j <= 10; j++ {
			fmt.Printf("%3d ", i*j)
		}
		// odřádkování
		fmt.Println()
	}
}
