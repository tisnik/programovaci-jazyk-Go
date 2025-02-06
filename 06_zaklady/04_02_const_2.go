// Základy programovacího jazyka Go
//
// - použití klíčového slova const
// - deklarace konstanty "a" globální pro celý modul
// - deklarace konstanty "b", která je platná jen v těle funkce "main"
// - deklarace konstanty "c", která je platná v jednom vnořeném bloku
// - deklarace konstanty "d", která je platná v dalším vnořeném bloku

package main

import "fmt"

// konstanta, která je globální pro celý modul
const a = 1

func main() {
	// konstanta, která je lokální v těle funkce
	const b = 2

	if a > 0 {
		const c = 3
		if a*b*c > 5 {
			const d = 4
			// na tomto místě můžeme použít všechny konstanty
			// - všechny jsou viditelné
			fmt.Println(a, b, c, d)
		}
		// zde není viditelná konstanta d,
		// ostatní tři konstanty ano
		fmt.Println(a, b, c)
	}
	// zde jsou viditelné jen konstanty a a b
	fmt.Println(a, b)
}
