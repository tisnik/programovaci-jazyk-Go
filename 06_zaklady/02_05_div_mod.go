// Základy programovacího jazyka Go
//
// - operátory pro celočíselný podíl a výpočet zbytku po dělení

package main

import "fmt"

// Výpočet podílu a zbytku po dělení
func computeDivMod(x, y int) {
	fmt.Printf("%3d / %2d = %3d   %3d %% %2d = %3d\n", x, y, x/y, x, y, x%y)
}

func main() {
	// otestování výpočtu podílu a zbytku po dělení pro všechny
	// kombinace kladných a záporných hodnot dělenců a dělitelů
	computeDivMod(10, 3)
	computeDivMod(-10, 3)
	computeDivMod(10, -3)
	computeDivMod(-10, -3)

	fmt.Println()

	// tabulka s výsledky zbytku po dělení 100/i pro i od 1 do 10
	for i := 1; i <= 10; i++ {
		computeDivMod(100, i)
	}
}
