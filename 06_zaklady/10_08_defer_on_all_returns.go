// Základy programovacího jazyka Go
//
// - ověření, zda se bude funkce s odloženým voláním
//   volat i v případě, že je použito větší množství
//   konstrukcí return

package main

import "fmt"

func function(x int) {
	fmt.Printf("Defer %d\n", x)
}

func classify(x int) string {
	// odložené volání funkce
	defer function(x)

	switch x {
	case 0:
		// (předčasný) návrat z funkce
		return "zero"
	case 2, 4, 6, 8:
		// (předčasný) návrat z funkce
		return "even number"
	case 1, 3, 5, 7, 9:
		// (předčasný) návrat z funkce
		return "odd number"
	default:
		// (předčasný) návrat z funkce
		return "?"
	}
}

func main() {
	// opakované volání funkce classify
	for x := 0; x <= 10; x++ {
		fmt.Println(x, classify(x))
	}
}
