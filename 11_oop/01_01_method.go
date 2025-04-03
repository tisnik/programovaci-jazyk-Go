// Metody v programovacím jazyku Go
//
// - deklarace záznamu
// - deklarace metody s příjemcem odpovídajícím záznamu
// - zavolání metody

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

func main() {
	// konstrukce obdélníka
	rect := Rectangle{width: 2, height: 3}

	// výpočet plochy
	a := rect.area()

	// zobrazení výsledků
	fmt.Printf("Rectangle %f x %f has area %f\n",
		rect.width, rect.height, a)
}
