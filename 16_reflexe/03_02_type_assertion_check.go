// Reflexe v programovacím jazyku Go
//
// - deklarace rozhraní s jedinou metodou
// - typová aserce
// - použití nekorektního rozhraní

package main

import (
	"fmt"
	"io"
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

// funkce akceptující libovolný typ splňující rozhraní io.Writer
// (POZOR: nikoli Shape)
func displayShapeInfo(shape io.Writer) {
	area := shape.area()
	fmt.Println("Area: ", area)

	// typová aserce
	square := shape.(Square)
	fmt.Println("Side length: ", square.sideLength)
}

func main() {
	s := Square{sideLength: 10}
	displayShapeInfo(s)
}
