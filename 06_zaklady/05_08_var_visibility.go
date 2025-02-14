// Základy programovacího jazyka Go
//
// - použití klíčového slova var
// - oblast viditelnosti proměnných

package main

import "fmt"

// globální proměnná
var x = 1

func main() {
	fmt.Println(x)

	// lokální proměnná pro funkci
	var x = 2
	fmt.Println(x)

	if true {
		// lokální proměnná pro blok
		var x = 3
		fmt.Println(x)

		// modifikace lokální proměnné
		x += 100
		fmt.Println(x)
	}
	// výpis lokální proměnné (v bloku se nezměnila)
	fmt.Println(x)

	// modifikace lokální proměnné (nikoli proměnné globální)
	x += 100
	fmt.Println(x)

	// tisk původní hodnoty globální proměnné
	printGlobalX()
}

func printGlobalX() {
	fmt.Println(x)
}
