// Základy programovacího jazyka Go
//
// - deklarace více uzávěrů, které sdílí společnou navázanou proměnnou
// - při konstrukci uzávěrů lze předat argumenty
//   s počáteční hodnotou a krokem

package main

import "fmt"

func newCounter(startValue, increment int) (func() int, func() int) {
	// interní proměnná použitá v uzávěru
	counter := startValue

	// interní funkce - uzávěr
	countUp := func() int {
		// přístup k vnější proměnné
		counter += increment
		return counter
	}

	// interní funkce - uzávěr
	countDown := func() int {
		// přístup k vnější proměnné
		counter -= increment
		return counter
	}
	// oba uzávěry jsou návratovými hodnotami
	// funkce newCounter
	return countUp, countDown
}

func main() {
	// konstrukce čítače nahoru i dolů s předáním
	// startovní hodnoty a kroku
	c_up, c_down := newCounter(100, 5)

	for range 5 {
		// změna čítače - počítání směrem nahoru
		fmt.Println(c_up())
	}

	for range 5 {
		// změna čítače - počítání směrem dolů
		fmt.Println(c_down())
	}
}
