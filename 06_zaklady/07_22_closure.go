// Základy programovacího jazyka Go
//
// - deklarace uzávěru
// - několik nezávislých čítačů

package main

import "fmt"

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
	// několik nezávislých čítačů
	counter1 := newCounter()
	counter2 := newCounter()
	counter3 := newCounter()
	counter4 := newCounter()

	// postupné zvyšování hodnot v nezávislých čítačích
	for range 10 {
		// opakované zavolání několika uzávěrů
		fmt.Println(counter1(), counter2(), counter3(), counter4())
	}
}
