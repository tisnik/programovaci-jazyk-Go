// Základy programovacího jazyka Go
//
// - deklarace uzávěru
// - konstruktor uzávěru
// - využití uzávěru z jiné funkce

package main

import "fmt"

// konstruktor čítače
func newCounter() func() int {
	// interní proměnná použitá v uzávěru
	counter := 0

	// interní funkce - uzávěr
	nextValue := func() int {
		// přístup k vnější proměnné
		counter++
		return counter
	}
	// uzávěr je návratovou hodnotou funkce newCounter
	return nextValue
}

func main() {
	// čítač
	counter := newCounter()

	// postupné zvyšování hodnoty čítačem
	for range 10 {
		// opakované zavolání uzávěru
		fmt.Println(counter())
	}
}
