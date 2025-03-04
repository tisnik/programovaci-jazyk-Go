// Základy programovacího jazyka Go
//
// - deklarace uzávěru
// - konstruktor uzávěru s parametry
// - využití uzávěru z jiné funkce

package main

import "fmt"

func newCounter(startValue, increment int) func() int {
	// interní proměnná použitá v uzávěru
	counter := startValue

	// interní funkce - uzávěr
	nextValue := func() int {
		// přístup k vnější proměnné
		counter += increment
		return counter
	}
	// uzávěr je návratovou hodnotou funkce newCounter
	return nextValue
}

func main() {
	// konstrukce čítače
	counter := newCounter(10, 2)

	for range 10 {
		// opakované zavolání uzávěru
		fmt.Println(counter())
	}
}
