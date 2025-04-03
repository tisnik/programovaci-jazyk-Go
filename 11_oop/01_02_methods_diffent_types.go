// Metody v programovacím jazyku Go
//
// - deklarace dvou nezávislých záznamů
// - deklarace metod s příjemcem odpovídajícím záznamu
// - zavolání metod

package main

import (
	"fmt"
)

// deklarace uživatelského typu
type Rectangle struct {
	width, height float64
}

// implementace metody
func (rectangle Rectangle) area() float64 {
	return rectangle.width * rectangle.height
}

// deklarace dalšího uživatelského typu
type Square struct {
	sideLength float64
}

// implementace metody
func (square Square) area() float64 {
	return square.sideLength * square.sideLength
}

func main() {
	// konstrukce čtverce
	square := Square{sideLength: 5}

	// konstrukce obdélníka
	rect := Rectangle{width: 2, height: 3}

	// výpočet plochy čtverce
	squareArea := square.area()

	// výpočet plochy obdélníku
	rectangleArea := rect.area()

	// zobrazení výsledků
	fmt.Printf("Square    %f x %f has area %f\n",
		square.sideLength, square.sideLength, squareArea)

	fmt.Printf("Rectangle %f x %f has area %f\n",
		rect.width, rect.height, rectangleArea)
}
