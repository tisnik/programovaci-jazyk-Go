// Základy programovacího jazyka Go
//
// - deklarace konstanty "a" globální pro celý modul
// - deklarace konstanty "b", která je platná jen v těle funkce "main"

package main

import "fmt"

// konstanta, která je globální pro celý modul
const a = 5

func main() {
	// konstanta, která je lokální v těle funkce
	const b = 6

	// na tomto místě můžeme použít obě konstanty - jsou viditelné
	fmt.Println(a, b)
}
