// Reflexe v programovacím jazyku Go
//
// - deklarace rozhraní s jedinou metodou
// - typová aserce

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

type Circle struct {
	radius float64
}

func (circle Circle) area() float64 {
	return math.Pi * circle.radius * circle.radius
}

// funkce akceptující libovolný typ splňující rozhraní Shape
func displayShapeInfo(shape Shape) {
	area := shape.area()
	fmt.Println("Area: ", area)

	square := shape.(Square)
	fmt.Println("Side length: ", square.sideLength)
}

func main() {
	s := Square{sideLength: 10}
	displayShapeInfo(s)

	r := Rectangle{width: 3, height: 2}
	displayShapeInfo(r)

	c := Circle{radius: 1}
	displayShapeInfo(c)
}
