// Základy programovacího jazyka Go
//
// - deklarace uzávěru
// - několik nezávislých čítačů

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
	// několik nezávislých čítačů
	counter1 := newCounter(0, 1)
	counter2 := newCounter(1, 1)
	counter3 := newCounter(0, 2)
	counter4 := newCounter(0, -1)

	// postupné zvyšování hodnot v nezávislých čítačích
	for range 10 {
		fmt.Println(counter1(), counter2(), counter3(), counter4())
	}
}
