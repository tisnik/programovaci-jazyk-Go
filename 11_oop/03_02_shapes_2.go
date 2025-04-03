// Koncept rozhraní v programovacím jazyku Go
//
// - deklarace rozhraní s jedinou metodou
// - deklarace uživatelského typu
// - funkce akceptující libovolný typ splňující rozhraní Shape
// - zavolání této funkce s hodnotou, která rozhraní Shape splňuje

package main

import (
	"fmt"
)

// deklarace rozhraní s jedinou metodou
type Shape interface {
	area() float64
}

// deklarace uživatelského typu
type Square struct {
	sideLength float64
}

// implementace metody předepsané v rozhraní Shape
func (square Square) area() float64 {
	return square.sideLength * square.sideLength
}

// funkce akceptující libovolný typ splňující rozhraní Shape
func displayArea(shape Shape) {
	area := shape.area()
	fmt.Println("Area: ", area)
}

func main() {
	s := Square{sideLength: 10}
	displayArea(s)
}
