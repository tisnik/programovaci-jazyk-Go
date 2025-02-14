// Základy programovacího jazyka Go
//
// - použití klíčového slova var
// - oblast viditelnosti proměnných
// - doplnění komentářů o hodnoty, které se vypíšou po spuštění

package main

import "fmt"

// globální proměnná
var x = 1

func main() {
	// vypíše se hodnota ** 1 **
	fmt.Println(x)

	// lokální proměnná pro funkci
	var x = 2

	// vypíše se hodnota ** 2 **
	fmt.Println(x)

	if true {
		// lokální proměnná pro blok
		var x = 3

		// vypíše se hodnota ** 3 **
		fmt.Println(x)

		// modifikace lokální proměnné
		x += 100

		// vypíše se hodnota ** 103 **
		fmt.Println(x)
	}
	// výpis lokální proměnné (v bloku se nezměnila)
	// vypíše se hodnota ** 2 **
	fmt.Println(x)

	// modifikace lokální proměnné (nikoli proměnné globální)
	x += 100

	// vypíše se hodnota ** 102 **
	fmt.Println(x)

	// tisk původní hodnoty globální proměnné
	// vypíše se hodnota ** 1 **
	printGlobalX()
}

func printGlobalX() {
	// vypíše se hodnota ** 1 **
	fmt.Println(x)
}
