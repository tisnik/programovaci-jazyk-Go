// Reflexe v programovacím jazyku Go
//
// - deklarace rozhraní s jedinou metodou
// - deklarace uživatelských typů splňujících stejné rozhraní
// - funkce akceptující libovolný typ splňující rozhraní Shape
// - zavolání této funkce s hodnotami, které rozhraní Shape splňují

package main

import (
	"fmt"
	"math"
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

// deklarace uživatelského typu
type Rectangle struct {
	width, height float64
}

// implementace metody předepsané v rozhraní Shape
func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

// deklarace uživatelského typu
type Circle struct {
	radius float64
}

// implementace metody předepsané v rozhraní Shape
func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// funkce akceptující libovolný typ splňující rozhraní Shape
func displayArea(shape Shape) {
	area := shape.area()
	fmt.Println("Area: ", area)
}

func main() {
	s := Square{sideLength: 10}
	displayArea(s)

	r := Rectangle{width: 3, height: 2}
	displayArea(r)

	c := Circle{radius: 1}
	displayArea(c)
}
