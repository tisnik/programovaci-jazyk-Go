// Základy programovacího jazyka Go
//
// - deklarace funkce s proměnným počtem parametrů
// - funkce je volána s různými počty argumentů

package main

import "fmt"

// Deklarace funkce s proměnným počtem parametrů, ale bez návratové hodnoty
func printSum(values ...int) {
	fmt.Printf("Called with %d arguments\n", len(values))
}

func main() {
	// zavolání funkce s předáním jednoho argumentu
	printSum(1)

	// zavolání funkce s předáním dvou argumentů
	printSum(1, 2)

	// zavolání funkce s předáním tří argumentů
	printSum(1, 2, 3)

	// zavolání funkce bez předání argumentů
	printSum()
}
