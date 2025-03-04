// Základy programovacího jazyka Go
//
// - deklarace více uzávěrů, které sdílí společnou navázanou proměnnou

package main

import "fmt"

func newCounter() (func() int, func() int) {
	// interní proměnná použitá v uzávěru
	counter := 0

	// interní funkce - uzávěr
	countUp := func() int {
		// přístup k vnější proměnné
		counter++
		return counter
	}

	// interní funkce - uzávěr
	countDown := func() int {
		// přístup k vnější proměnné
		counter--
		return counter
	}
	// oba uzávěry jsou návratovými hodnotami
	// funkce newCounter
	return countUp, countDown
}

func main() {
	// čítač nahoru i dolů
	c_up, c_down := newCounter()

	for range 5 {
		// změna čítače - počítání směrem nahoru
		fmt.Println(c_up())
	}

	for range 5 {
		// změna čítače - počítání směrem dolů
		fmt.Println(c_down())
	}
}
